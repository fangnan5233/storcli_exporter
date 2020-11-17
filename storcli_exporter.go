package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	argStorcliPath   = flag.String("storcli_path", "/opt/MegaRAID/storcli/storcli64", "Path to MegaRAID StorCLI or PercCLI binary. By default '/opt/MegaRAID/storcli/storcli64'.")
	argMetricsPath   = flag.String("metrics_path", "/metrics", "Path under which to expose Prometheus metrics. By default '/metrics'.")
	argMetricsPrefix = flag.String("metrics_prefix", "storcli", "Prefix for Prometheus metrics. By default 'storcli'.")
	argListenAddress = flag.String("listen_address", ":9326", "Listen address for this exporter. By default ':9326'.")
)

type Response struct {
	Controllers []struct {
		ResponseData struct {
			VirtualDrives int `json:"Virtual Drives"`
			VDLIST        []struct {
				Position string `json:"DG/VD"`
				Type     string `json:"TYPE"`
				State    string `json:"State"`
				Size     string `json:"Size"`
			} `json:"VD LIST"`
			PhysicalDrives int `json:"Physical Drives"`
			PDLIST         []struct {
				Device   int    `json:"DID"`
				Position string `json:"EID:Slt"`
				State    string `json:"State"`
				Media    string `json:"Med"`
				Model    string `json:"Model"`
				Size     string `json:"Size"`
				Type     string `json:"Type"`
			} `json:"PD LIST"`
		} `json:"Response Data"`
	} `json:"Controllers"`
}

type Exporter struct {
	physicalDriveStatus *prometheus.Desc
	virtualDriveStatus  *prometheus.Desc
	physicalDriveCount  *prometheus.Desc
	virtualDriveCount   *prometheus.Desc
	scrapeSuccess       *prometheus.Desc
}

func fetchStorcliOutput() (resp Response, err error) {
	output, err := exec.Command(*argStorcliPath, "/call", "show", "all", "J").Output()
	if err != nil {
		return Response{}, fmt.Errorf("Failed to execute command: %s", err)
	}
	var response Response
	err = json.Unmarshal(output, &response)
	if err != nil {
		return Response{}, fmt.Errorf("Failed to unmarshal JSON: %s", err)
	}
	return response, nil
}

func NewExporter() *Exporter {
	return &Exporter{
		scrapeSuccess:       ScrapeSuccess,
		virtualDriveCount:   VirtualDrivesCount,
		physicalDriveCount:  PhysicalDrivesCount,
		virtualDriveStatus:  VirtualDriveStatus,
		physicalDriveStatus: PhysicalDriveStatus,
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	ch <- e.physicalDriveStatus
	ch <- e.virtualDriveStatus
	ch <- e.physicalDriveCount
	ch <- e.virtualDriveCount
	ch <- e.scrapeSuccess
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	response, err := fetchStorcliOutput()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(e.scrapeSuccess, prometheus.GaugeValue, 0)
		log.Printf("Failed to fetch StorCLI output: %s", err)
	}
	for controllerNumber, controller := range response.Controllers {
		ch <- prometheus.MustNewConstMetric(e.virtualDriveCount, prometheus.GaugeValue, float64(controller.ResponseData.VirtualDrives), strconv.Itoa(controllerNumber))
		ch <- prometheus.MustNewConstMetric(e.physicalDriveCount, prometheus.GaugeValue, float64(controller.ResponseData.PhysicalDrives), strconv.Itoa(controllerNumber))
		for _, virtualDrive := range controller.ResponseData.VDLIST {
			ch <- prometheus.MustNewConstMetric(
				e.virtualDriveStatus, prometheus.GaugeValue, 1.0,
				strconv.Itoa(controllerNumber), virtualDrive.Position, virtualDrive.Type, virtualDrive.Size, virtualDrive.State,
			)
		}
		for _, physicalDrive := range controller.ResponseData.PDLIST {
			ch <- prometheus.MustNewConstMetric(
				e.physicalDriveStatus, prometheus.GaugeValue, 1.0,
				strconv.Itoa(controllerNumber), physicalDrive.Position, strconv.Itoa(physicalDrive.Device), physicalDrive.Model,
				physicalDrive.State, physicalDrive.Media, physicalDrive.Size,
			)
		}
	}
}

func main() {
	flag.Parse()

	registry := prometheus.NewRegistry()
	registry.MustRegister(NewExporter())

	http.Handle(*argMetricsPath, promhttp.HandlerFor(registry, promhttp.HandlerOpts{}))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		_, err := w.Write([]byte(`<html>
    <head><title>StorCLI Exporter</title></head>
    <body>
    <h1>StorCLI Exporter</h1>
    <p><a href='` + *argMetricsPath + `'>Metrics</a></p>
    </html>`))
		if err != nil {
			log.Printf("Failed to write to HTTP client: %s", err)
		}
	})

	server := &http.Server{
		Addr:         *argListenAddress,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
	}

	log.Printf("StorCLI exporter started and listening on %s", server.Addr)
	log.Fatal(server.ListenAndServe())
}
