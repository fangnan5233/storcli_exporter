package main

import "github.com/prometheus/client_golang/prometheus"

var (
	// ScrapeSuccess returns new Prometheus metric description
	ScrapeSuccess = prometheus.NewDesc(
		prometheus.BuildFQName(*argMetricsPrefix, "", "scrape_success"),
		"Was the last scrape of StorCLI successfull.",
		nil,
		nil,
	)
	// VirtualDrivesCount returns new Prometheus metric description
	VirtualDrivesCount = prometheus.NewDesc(
		prometheus.BuildFQName(*argMetricsPrefix, "", "virtual_drive_count"),
		"Count of available Virtual Drives.",
		[]string{"controller"},
		nil,
	)
	// PhysicalDrivesCount returns new Prometheus metric description
	PhysicalDrivesCount = prometheus.NewDesc(
		prometheus.BuildFQName(*argMetricsPrefix, "", "physical_drive_count"),
		"Count of available Physical Drives.",
		[]string{"controller"},
		nil,
	)
	// VirtualDriveStatus returns new Prometheus metric description
	VirtualDriveStatus = prometheus.NewDesc(
		prometheus.BuildFQName(*argMetricsPrefix, "", "virtual_drive_status"),
		"Status of the Virtual Drive.",
		[]string{
			"controller", "slot", "type", "size", "state",
		},
		nil,
	)
	// PhysicalDriveStatus returns new Prometheus metric description
	PhysicalDriveStatus = prometheus.NewDesc(
		prometheus.BuildFQName(*argMetricsPrefix, "", "physical_drive_status"),
		"Status of the Physical Drive.",
		[]string{
			"controller", "slot", "device", "model",
			"state", "media", "size",
		},
		nil,
	)
)
