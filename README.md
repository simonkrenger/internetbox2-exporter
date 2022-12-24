# internetbox2-exporter

A Prometheus exporter to read metrics from a Swisscom Internet-Box 2 router.

Under the hood, the device is an [Askey RTV1905VW](https://deviwiki.com/wiki/Askey_RTV1905VW) (SG2-NP-00) with a custom interface.

## Running the exporter

```
$ podman run --rm -d -p 8088:8088 -e INTERNETBOX_ADMIN_PASSWORD="secret" quay.io/simonkrenger/internetbox2-exporter:latest
$ curl localhost:8088/metrics
...
# HELP internetbox_client_devices_count Currently connected devices, by interface
# TYPE internetbox_client_devices_count gauge
internetbox_client_devices_count 6
...
```

If you cannot connect to `http://internetbox.home` from the container, you may need to set the environment variable `-e INTERNETBOX_BASEURL="http://192.168.0.1"`.

The exporter exports the following metrics:

* internetbox_client_devices_count
* internetbox_client_noise
* internetbox_client_signalstrength
* internetbox_client_total_rx_bytes
* internetbox_client_total_tx_bytes
* internetbox_device_load_avg_1
* internetbox_device_load_avg_5
* internetbox_device_load_avg_15
* internetbox_device_memory_free_bytes
* internetbox_device_memory_total_bytes
* internetbox_device_memory_usage_bytes
* internetbox_device_processes
* internetbox_device_reboots
* internetbox_device_swap_free_bytes
* internetbox_device_swap_total_bytes
* internetbox_device_temperature_celsius
* internetbox_device_uptime_seconds
* internetbox_dsl_errors_received_total
* internetbox_dsl_errors_sent_total
* internetbox_dsl_packets_received_total
* internetbox_dsl_packets_sent_total
* internetbox_dsl_received_bytes
* internetbox_dsl_sent_bytes
* internetbox_wifi_air_use_percentage
* internetbox_wifi_band_total_received_bytes
* internetbox_wifi_band_total_sent_bytes
* internetbox_wifi_data_mbps
* internetbox_wifi_data_use_percentage
* internetbox_wifi_phy_mbps
* internetbox_wifi_signal_noise_ratio
* internetbox_wifi_signal_strength_db
* internetbox_wifi_total_received_bytes
* internetbox_wifi_total_sent_bytes

## Internet-Box 2 Endpoints

When capturing the requests that the web interface makes, we can see calls to the following endpoints. Many of these endpoints are available while unauthenticated, some return more data when authenticated. The below table shows the endpoints used by this Prometheus exporter.

| Path |  Description |
|---|---|
| /ws | Multi-purpose endpoint for **login**, **event management**, **device overview** |
| /sysbus/APController/LAN:get | TODO |
| /sysbus/APController/PowerManagement:getMode | TODO |
| /sysbus/APController:get | TODO |
| /sysbus/DECT:getBaseState | TODO |
| /sysbus/DECT:getNEMoState | TODO |
| /sysbus/DECT:getRFPI | TODO |
| /sysbus/DECT:getVersion | TODO |
| /sysbus/DeviceInfo:get | Endpoint to retrieve common device information such as manufacturer, model name and software version |
| /sysbus/DeviceInfo/MemoryStatus:get | TODO |
| /sysbus/DHCPv4/Server/Pool/default:getStaticLeases | TODO |
| /sysbus/DHCPv4/Server:getDHCPServerPool | TODO |
| /sysbus/DHCPv6/Server:getPrefixInformation | TODO |
| /sysbus/DNS/Server/Cache:get | TODO |
| /sysbus/DNS/Server/Route:get | TODO |
| /sysbus/DNS:get | TODO |
| /sysbus/DSPGDECT:get | TODO |
| /sysbus/DynDNS:getGlobalEnable | TODO |
| /sysbus/DynDNS:getHosts | TODO |
| /sysbus/Firewall:getCustomRule | TODO |
| /sysbus/Firewall:getDMZ | TODO |
| /sysbus/Firewall:getFirewallIPv6Level | TODO |
| /sysbus/Firewall:getPortForwarding | TODO |
| /sysbus/GenLog:nextLogs | Endpoint to get log data from the device |
| /sysbus/MultipathTCP:get | TODO |
| /sysbus/NeMo/Intf/bridge:getFirstParameter | TODO |
| /sysbus/NeMo/Intf/data:getMIBs | TODO |
| /sysbus/NeMo/Intf/data:hasFlag | TODO |
| /sysbus/NeMo/Intf/dsl0:getDSLChannelStats | Endpoint to read DSL channel statistics such as bytes sent and received. |
| /sysbus/NeMo/Intf/guest:getMIBs | TODO |
| /sysbus/NeMo/Intf/lan:getFirstParameter | TODO |
| /sysbus/NeMo/Intf/lan:getMIBs | TODO |
| /sysbus/NeMo/Intf/wifi0_bcm:get | TODO |
| /sysbus/NeMo/Intf/wifi0_bcm:getLatestPower | TODO |
| /sysbus/NeMo/Intf/wifi0_bcm:getPerAntennaRssi | TODO |
| /sysbus/NeMo/Intf/wifi1_bcm:get | TODO |
| /sysbus/NeMo/Intf/wifi1_bcm:getLatestPower | TODO |
| /sysbus/NeMo/Intf/wifi1_bcm:getPerAntennaRssi | TODO |
| /sysbus/NeMo/Intf/wl0:getStationStats | TODO |
| /sysbus/NeMo/Intf/wl1:getStationStats | TODO |
| /sysbus/NeMo/Intf/wwan:getMIBs | TODO |
| /sysbus/NetMaster/LAN/default/Bridge/lan:getIPv4 | TODO |
| /sysbus/NetMaster/LAN/default/Bridge/lan:getIPv6Configuration | TODO |
| /sysbus/NetMaster:get | TODO |
| /sysbus/NetWWAN/Device/wwan_xt:getResults | TODO |
| /sysbus/NetWWAN/Device:get | TODO |
| /sysbus/NMC/AutoIPv6:get | TODO |
| /sysbus/NMC/Debug:getWifiDriverStats | Endpoint to get WiFi driver statistics |
| /sysbus/NMC/Guest:get | TODO |
| /sysbus/NMC/LAN:getStaticRoutes | TODO |
| /sysbus/NMC/LEDs/LED/Status:get | TODO |
| /sysbus/NMC/LEDs/LED/Wireless:get | TODO |
| /sysbus/NMC/PhonebookBackup:get | TODO |
| /sysbus/NMC/Wifi:get | TODO |
| /sysbus/NMC/WWAN:get | TODO |
| /sysbus/NMC:get | TODO |
| /sysbus/NMC:getLANIP | Retrieve IP information about the LAN ports |
| /sysbus/NMC:getWANStatus | Retrieve status of the WAN Port |
| /sysbus/NMC/WWAN:get | TODO |
| /sysbus/OopsTracker:get | TODO |
| /sysbus/OopsTracker:getOopses | TODO |
| /sysbus/PasswordRecovery:getStatus | TODO |
| /sysbus/Scheduler:getCompleteSchedules | TODO |
| /sysbus/SSW/Steering/MasterConfig/OperatingStandards:get | TODO |
| /sysbus/SSW/Steering/MasterConfig:get | TODO |
| /sysbus/SSW/Steering/SteerLog/Attempt:get | TODO |
| /sysbus/SSW/Steering:get | TODO |
| /sysbus/Time:getTime | TODO |
| /sysbus/USBHosts:getDevices | TODO |
| /sysbus/UserManagement:getUserLog | TODO |
| /sysbus/VoiceService/VoiceApplication:getSipExtensionsStatus | TODO |
| /sysbus/VoiceService/VoiceApplication:listGroups | TODO |
| /sysbus/VoiceService/VoiceApplication:listHandsets | TODO |
| /sysbus/VoiceService/VoiceApplication:listTrunks | TODO |
| /sysbus/VPN:get | TODO |
| /sysbus/VPN:getServersConfig | TODO |
| /sysbus/WebuiupgradeService:get | TODO |
| /sysbus/WebuiupgradeService:getLatestVersion | TODO |
