# StorCLI Exporter
Export Prometheus metrics for StorCLI and PercCLI RAID utilities

## Supported MegaRAID utilities
* PercCLI (Basically copy of StorCLI)
* StorCLI (Tested with version 1.20.15)

## Status
**Production (Released)**

## Up && Running
Download the executable binary and run it with
```
./storcli_exporter <flags>
```

## Flags
```
./storcli_exporter -h
Usage of ./storcli_exporter:
  -listen_address string
    	Listen address for this exporter. By default ':9326'. (default ":9326")
  -metrics_path string
    	Path under which to expose Prometheus metrics. By default '/metrics'. (default "/metrics")
  -metrics_prefix string
    	Prefix for Prometheus metrics. By default 'storcli'. (default "storcli")
  -storcli_path string
    	Path to MegaRAID StorCLI or PercCLI binary. By default '/opt/MegaRAID/storcli/storcli64'. (default "/opt/MegaRAID/storcli/storcli64")
```

## Prometheus
Exporter metrics for Prometheus
```
# HELP storcli_physical_drive_count Count of available Physical Drives.
# TYPE storcli_physical_drive_count gauge
storcli_physical_drive_count{controller="0"} 6
# HELP storcli_physical_drive_status Status of the Physical Drive.
# TYPE storcli_physical_drive_status gauge
storcli_physical_drive_status{controller="0",device="0",enclosure="252",media="HDD",model="ST4000NM0035-1V4107",size="3.637 TB",slot="0",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="1",enclosure="252",media="HDD",model="ST4000NM0035-1V4107",size="3.637 TB",slot="1",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="4",enclosure="252",media="SSD",model="Samsung SSD 860 PRO 1TB",size="953.343 GB",slot="5",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="5",enclosure="252",media="SSD",model="Samsung SSD 860 PRO 1TB",size="953.343 GB",slot="4",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="6",enclosure="252",media="HDD",model="ST4000NM0035-1V4107",size="3.637 TB",slot="2",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="7",enclosure="252",media="HDD",model="ST4000NM0035-1V4107",size="3.637 TB",slot="3",state="Onln"} 1
# HELP storcli_virtual_drive_count Count of available Virtual Drives.
# TYPE storcli_virtual_drive_count gauge
storcli_virtual_drive_count{controller="0"} 2
# HELP storcli_virtual_drive_status Status of the Virtual Drive.
# TYPE storcli_virtual_drive_status gauge
storcli_virtual_drive_status{controller="0",size="7.276 TB",slot="1/1",state="Optl",type="RAID10"} 1
storcli_virtual_drive_status{controller="0",size="953.343 GB",slot="0/0",state="Optl",type="RAID1"} 1
```

## Contributing
Pull Requests are welcome and appreciated. For more major changes, create an issue in this repository to discuss the changes before creating a Pull Request.

## License
Distributed under the MIT License. See LICENSE for more information.

## Contact
Konradas Bunikis - konradas.bunikis@zohomail.eu
