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
# HELP storcli_drive_groups_count Count of available Drive Groups.
# TYPE storcli_drive_groups_count gauge
storcli_drive_groups_count{controller="0"} 3
# HELP storcli_physical_drive_count Count of available Physical Drives.
# TYPE storcli_physical_drive_count gauge
storcli_physical_drive_count{controller="0"} 14
# HELP storcli_physical_drive_status Status of the Physical Drive.
# TYPE storcli_physical_drive_status gauge
storcli_physical_drive_status{controller="0",device="1",media="SSD",model="INTEL SSDSC2KB038T8",size="3.492 TB",slot="0:2",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="10",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:11",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="11",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:7",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="12",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:5",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="13",media="HDD",model="ST8000NM000A-2KE",size="7.276 TB",slot="0:9",state="UBad"} 1
storcli_physical_drive_status{controller="0",device="14",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:13",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="2",media="SSD",model="INTEL SSDSC2KB038T8",size="3.492 TB",slot="0:3",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="3",media="SSD",model="SAMSUNG MZ7LH480HAHQ-00005",size="446.625 GB",slot="0:0",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="4",media="SSD",model="SAMSUNG MZ7LH480HAHQ-00005",size="446.625 GB",slot="0:1",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="5",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:4",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="6",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:6",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="7",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:8",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="8",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:12",state="Onln"} 1
storcli_physical_drive_status{controller="0",device="9",media="HDD",model="ST8000NM000A-2KE101",size="7.276 TB",slot="0:10",state="Onln"} 1
# HELP storcli_topology_status Status of the opology.
# TYPE storcli_topology_status gauge
storcli_topology_status{array="-",controller="0",device="-",disk_group="0",row="-",size="446.625 GB",slot="-",state="Optl",type="RAID1"} 1
storcli_topology_status{array="-",controller="0",device="-",disk_group="1",row="-",size="6.984 TB",slot="-",state="Optl",type="RAID0"} 1
storcli_topology_status{array="-",controller="0",device="-",disk_group="2",row="-",size="65.491 TB",slot="-",state="Dgrd",type="RAID5"} 1
storcli_topology_status{array="0",controller="0",device="-",disk_group="0",row="-",size="446.625 GB",slot="-",state="Optl",type="RAID1"} 1
storcli_topology_status{array="0",controller="0",device="-",disk_group="1",row="-",size="6.984 TB",slot="-",state="Optl",type="RAID0"} 1
storcli_topology_status{array="0",controller="0",device="-",disk_group="2",row="-",size="65.491 TB",slot="-",state="Dgrd",type="RAID5"} 1
storcli_topology_status{array="0",controller="0",device="-",disk_group="2",row="8",size="7.276 TB",slot="-",state="Msng",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="1",disk_group="1",row="0",size="3.492 TB",slot="0:2",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="10",disk_group="2",row="5",size="7.276 TB",slot="0:11",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="11",disk_group="2",row="6",size="7.276 TB",slot="0:7",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="12",disk_group="2",row="7",size="7.276 TB",slot="0:5",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="14",disk_group="2",row="9",size="7.276 TB",slot="0:13",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="2",disk_group="1",row="1",size="3.492 TB",slot="0:3",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="3",disk_group="0",row="0",size="446.625 GB",slot="0:0",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="4",disk_group="0",row="1",size="446.625 GB",slot="0:1",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="5",disk_group="2",row="0",size="7.276 TB",slot="0:4",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="6",disk_group="2",row="1",size="7.276 TB",slot="0:6",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="7",disk_group="2",row="2",size="7.276 TB",slot="0:8",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="8",disk_group="2",row="3",size="7.276 TB",slot="0:12",state="Onln",type="DRIVE"} 1
storcli_topology_status{array="0",controller="0",device="9",disk_group="2",row="4",size="7.276 TB",slot="0:10",state="Onln",type="DRIVE"} 1
# HELP storcli_virtual_drive_count Count of available Virtual Drives.
# TYPE storcli_virtual_drive_count gauge
storcli_virtual_drive_count{controller="0"} 3
# HELP storcli_virtual_drive_status Status of the Virtual Drive.
# TYPE storcli_virtual_drive_status gauge
storcli_virtual_drive_status{controller="0",size="446.625 GB",slot="0/0",state="Optl",type="RAID1"} 1
storcli_virtual_drive_status{controller="0",size="6.984 TB",slot="1/1",state="Optl",type="RAID0"} 1
storcli_virtual_drive_status{controller="0",size="65.491 TB",slot="2/2",state="Dgrd",type="RAID5"} 1
```

## Contributing
Pull Requests are welcome and appreciated. For more major changes, create an issue in this repository to discuss the changes before creating a Pull Request.

## License
Distributed under the MIT License. See LICENSE for more information.

## Contact
Konradas Bunikis - konradas.bunikis@zohomail.eu
