package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
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
				DGVD  string `json:"DG/VD"`
				TYPE  string `json:"TYPE"`
				State string `json:"State"`
				Size  string `json:"Size"`
			} `json:"VD LIST"`
			PhysicalDrives int `json:"Physical Drives"`
			PDLIST         []struct {
				EIDSlt string `json:"EID:Slt"`
				DID    int    `json:"DID"`
				State  string `json:"State"`
				Size   string `json:"Size"`
				Med    string `json:"Med"`
				Model  string `json:"Model"`
				Type   string `json:"Type"`
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
		scrapeSuccess:       prometheus.NewDesc(prometheus.BuildFQName(*argMetricsPrefix, "", "scrape_success"), "Was the last scrape of StorCLI successfull.", nil, nil),
		virtualDriveCount:   prometheus.NewDesc(prometheus.BuildFQName(*argMetricsPrefix, "", "virtual_drive_count"), "Count of available Virtual Drives.", []string{"controller"}, nil),
		physicalDriveCount:  prometheus.NewDesc(prometheus.BuildFQName(*argMetricsPrefix, "", "physical_drive_count"), "Count of available Physical Drives.", []string{"controller"}, nil),
		virtualDriveStatus:  prometheus.NewDesc(prometheus.BuildFQName(*argMetricsPrefix, "", "virtual_drive_status"), "Status of the Virtual Drive.", []string{"controller", "slot", "type", "size", "state"}, nil),
		physicalDriveStatus: prometheus.NewDesc(prometheus.BuildFQName(*argMetricsPrefix, "", "physical_drive_status"), "Status of the Physical Drive.", []string{"controller", "enclosure", "device", "model", "state", "media", "slot", "size"}, nil),
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
			var value float64
			if virtualDrive.State == "Optl" {
				value = 1.0
			} else {
				value = 0.0
			}
			ch <- prometheus.MustNewConstMetric(e.virtualDriveStatus, prometheus.GaugeValue, value, strconv.Itoa(controllerNumber), virtualDrive.DGVD, virtualDrive.TYPE, virtualDrive.Size, virtualDrive.State)
		}
		for _, physicalDrive := range controller.ResponseData.PDLIST {
			eidslt := strings.Split(physicalDrive.EIDSlt, ":")
			enclosure, slot := eidslt[0], eidslt[1]
			var value float64
			if physicalDrive.State == "Onln" {
				value = 1.0
			} else {
				value = 0.0
			}
			ch <- prometheus.MustNewConstMetric(e.physicalDriveStatus, prometheus.GaugeValue, value, strconv.Itoa(controllerNumber), enclosure, strconv.Itoa(physicalDrive.DID), physicalDrive.Model, physicalDrive.State, physicalDrive.Med, slot, physicalDrive.Size)
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
