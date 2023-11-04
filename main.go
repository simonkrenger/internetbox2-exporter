package main

import (
	"internetbox2-golang/sysbus"
	"internetbox2-golang/sysbus/nemo"
	"internetbox2-golang/sysbus/nmc"
	"internetbox2-golang/ws"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	// Record metrics every minute
	go func() {
		for {
			// Device information from /sysbus/DeviceInfo:get
			deviceInfo, err := sysbus.GetDeviceInfo()
			if err != nil {
				log.Fatalln(err)
			}
			ibDeviceUptime.Set(float64(deviceInfo.Status.UpTime))
			ibDeviceNumberOfReboots.Set(float64(deviceInfo.Status.NumberOfReboots))

			// More device information, this time from another endpoint
			routerEvents, err := ws.GetInternetBoxStatus()
			if err != nil || len(routerEvents.Status.Events) == 0 {
				log.Fatalln(err)
			}
			var deviceStats = routerEvents.Status.Events[0].Data.Object.Attributes.Hgw
			// Load stats
			ibDeviceLoadAvg1Min.Set(float64(deviceStats.Loadaverage1Min))
			ibDeviceLoadAvg5Min.Set(float64(deviceStats.Loadaverage5Min))
			ibDeviceLoadAvg15Min.Set(float64(deviceStats.Loadaverage15Min))
			ibDeviceNumberOfProcesses.Set(float64(deviceStats.Procs))
			// Memory stats
			ibDeviceMemoryTotal.Set(float64(deviceStats.Totalram))
			ibDeviceMemoryFree.Set(float64(deviceStats.Freeram))
			ibDeviceSwapTotal.Set(float64(deviceStats.Totalswap))
			ibDeviceSwapFree.Set(float64(deviceStats.Freeswap))
			ibDevicMemoryUsage.With(prometheus.Labels{"type": "buffer"}).Set(float64(deviceStats.Bufferram))
			ibDevicMemoryUsage.With(prometheus.Labels{"type": "cached"}).Set(float64(deviceStats.Cachedram))
			ibDevicMemoryUsage.With(prometheus.Labels{"type": "shared"}).Set(float64(deviceStats.Sharedram))

			// DSL channel stats
			DSLChannelStats, err := nemo.GetDSLChannelStats("dsl0")
			if err != nil {
				log.Fatalln(err)
			}
			ibDSLChannelErrorsSent.Set(float64(DSLChannelStats.Status.Stats.ErrorsSent))
			ibDSLChannelErrorsReceived.Set(float64(DSLChannelStats.Status.Stats.ErrorsReceived))
			ibDSLChannelBytesSent.Set(float64(DSLChannelStats.Status.Stats.BytesSent))
			ibDSLChannelBytesReceived.Set(float64(DSLChannelStats.Status.Stats.BytesReceived))
			ibDSLChannelPacketsSent.Set(float64(DSLChannelStats.Status.Stats.PacketsSent))
			ibDSLChannelPacketsReceived.Set(float64(DSLChannelStats.Status.Stats.PacketsReceived))

			// WiFi stats
			for _, wifiInterface := range []string{"wl0", "wl1"} {
				stationStats, err := nemo.GetWifiStationStats(wifiInterface)
				if err != nil {
					log.Fatalln(err)
				}
				for _, station := range stationStats.Status {
					ibWifiBytesReceived.With(prometheus.Labels{"channel": wifiInterface, "station": station.MACAddress}).Set(float64(station.RxBytes))
					ibWifiBytesSent.With(prometheus.Labels{"channel": wifiInterface, "station": station.MACAddress}).Set(float64(station.TxBytes))
					ibWifiSignalStrength.With(prometheus.Labels{"channel": wifiInterface, "station": station.MACAddress}).Set(float64(station.SignalStrength))
					ibWifiSignalNoiseRatio.With(prometheus.Labels{"channel": wifiInterface, "station": station.MACAddress}).Set(float64(station.SignalNoiseRatio))
				}
			}

			// More WiFi stats
			wifiStats, err := nmc.GetWifiDriverStats()
			if err != nil {
				log.Fatalln(err)
			}
			ibWifiBytesSentByBand.With(prometheus.Labels{"band": "2.4GHz"}).Set(float64(wifiStats.Status.Band24G.Counters.Txbyte))
			ibWifiBytesReceivedByBand.With(prometheus.Labels{"band": "2.4GHz"}).Set(float64(wifiStats.Status.Band24G.Counters.Rxbyte))
			ibWifiBytesSentByBand.With(prometheus.Labels{"band": "5GHz"}).Set(float64(wifiStats.Status.Band5G.Counters.Txbyte))
			ibWifiBytesReceivedByBand.With(prometheus.Labels{"band": "5GHz"}).Set(float64(wifiStats.Status.Band5G.Counters.Rxbyte))

			ibDeviceChipTemperature.With(prometheus.Labels{"band": "2.4GHz"}).Set(float64(nmc.ParseTemperature(wifiStats.Status.Band24G.Parameters.Temperature)))
			ibDeviceChipTemperature.With(prometheus.Labels{"band": "5GHz"}).Set(float64(nmc.ParseTemperature(wifiStats.Status.Band5G.Parameters.Temperature)))

			for _, station := range wifiStats.Status.Band24G.Stations {
				ibWifiDataMbps.With(prometheus.Labels{"band": "2.4GHz", "station": station.Mac}).Set(nmc.ParseMbps(station.DataMbps))
				ibWifiPhyMbps.With(prometheus.Labels{"band": "2.4GHz", "station": station.Mac}).Set(nmc.ParseMbps(station.PHYMbps))
				ibWifiAirUse.With(prometheus.Labels{"band": "2.4GHz", "station": station.Mac}).Set(nmc.ParsePercentage(station.AirUse))
				ibWifiDataUse.With(prometheus.Labels{"band": "2.4GHz", "station": station.Mac}).Set(nmc.ParsePercentage(station.DataUse))
			}
			for _, station := range wifiStats.Status.Band5G.Stations {
				ibWifiDataMbps.With(prometheus.Labels{"band": "5GHz", "station": station.Mac}).Set(nmc.ParseMbps(station.DataMbps))
				ibWifiPhyMbps.With(prometheus.Labels{"band": "5GHz", "station": station.Mac}).Set(nmc.ParseMbps(station.PHYMbps))
				ibWifiAirUse.With(prometheus.Labels{"band": "5GHz", "station": station.Mac}).Set(nmc.ParsePercentage(station.AirUse))
				ibWifiDataUse.With(prometheus.Labels{"band": "5GHz", "station": station.Mac}).Set(nmc.ParsePercentage(station.DataUse))
			}

			// Connected device statistics
			lanDevices, err := ws.GetDevices("lan and not self and not interface")
			if err != nil {
				log.Fatalln(err)
			}
			noOfActiveDevices := 0
			for _, d := range lanDevices.Status {
				if d.Active {
					noOfActiveDevices++
					switch d.Layer2Interface {
					case "WL1":

						events, _ := ws.GetEventLogForDevice(d.PhysAddress)
						if err != nil {
							log.Println(err)
							continue
						}
						lastEvent := len(events.Status) - 1
						if lastEvent <= 0 {
							continue
						}
						ibClientRxBytes.With(prometheus.Labels{"station": d.PhysAddress, "ip": d.IPAddress, "interface": d.InterfaceName}).Set(float64(events.Status[lastEvent].RxBytes))
						ibClientTxBytes.With(prometheus.Labels{"station": d.PhysAddress, "ip": d.IPAddress, "interface": d.InterfaceName}).Set(float64(events.Status[lastEvent].TxBytes))
						ibClientSignalStrength.With(prometheus.Labels{"station": d.PhysAddress, "ip": d.IPAddress, "interface": d.InterfaceName}).Set(float64(events.Status[lastEvent].SignalStrength))
						ibClientNoise.With(prometheus.Labels{"station": d.PhysAddress, "ip": d.IPAddress, "interface": d.InterfaceName}).Set(float64(events.Status[lastEvent].Noise))
					case "eth1.0":
						log.Println("Layer 2 type for " + d.PhysAddress + " is 'eth1.0', does not have any bandwidth information. Skipping.")
					default:
						log.Println("Unknown layer 2 type for " + d.PhysAddress + ", skipping: " + d.Layer2Interface)
					}
				}
			}
			ibClientsNumberOfActiveDevices.Set(float64(noOfActiveDevices))

			// Fixed to 1 minute interval
			time.Sleep(60 * time.Second)
		}
	}()
}

func main() {

	adminPass := os.Getenv("INTERNETBOX_ADMIN_PASSWORD")
	if adminPass == "" {
		log.Fatalln("INTERNETBOX_ADMIN_PASSWORD not set")
		os.Exit(1)
	}

	// Start actual gathering
	recordMetrics()

	// Publish prometheus endpoints
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8088", nil)

}
