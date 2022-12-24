package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Device metrics
	ibDeviceUptime = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_uptime_seconds",
		Help: "Uptime of the device in seconds",
	})
	ibDeviceNumberOfReboots = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_reboots",
		Help: "Number of reboots of the device",
	})
	ibDeviceLoadAvg1Min = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_load_avg_1",
		Help: "1 minute average for device load",
	})
	ibDeviceLoadAvg5Min = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_load_avg_5",
		Help: "5 minute average for device load",
	})
	ibDeviceLoadAvg15Min = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_load_avg_15",
		Help: "15 minute average for device load",
	})
	ibDeviceNumberOfProcesses = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_processes",
		Help: "Number of processes",
	})
	ibDeviceChipTemperature = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_device_temperature_celsius",
		Help: "Temperature of the WiFi chip",
	}, []string{"band"})
	ibDeviceMemoryTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_memory_total_bytes",
		Help: "Total RAM of the device",
	})
	ibDeviceMemoryFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_memory_free_bytes",
		Help: "Free memory of the device",
	})
	ibDevicMemoryUsage = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_device_memory_usage_bytes",
		Help: "Used memory on the device, by type",
	}, []string{"type"})
	ibDeviceSwapTotal = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_swap_total_bytes",
		Help: "Total swap space on the device",
	})
	ibDeviceSwapFree = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_device_swap_free_bytes",
		Help: "Free swap space on the device",
	})

	// DSL metrics
	ibDSLChannelErrorsSent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_dsl_errors_sent_total",
		Help: "Number of errors sent on the DSL channel",
	})
	ibDSLChannelErrorsReceived = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_dsl_errors_received_total",
		Help: "Number of errors received on the DSL channel",
	})
	ibDSLChannelBytesSent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_dsl_sent_bytes",
		Help: "Bytes sent on the DSL channel",
	})
	ibDSLChannelBytesReceived = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_dsl_received_bytes",
		Help: "Bytes received on the DSL channel",
	})
	ibDSLChannelPacketsSent = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_dsl_packets_sent_total",
		Help: "Number of packets sent on the DSL channel",
	})
	ibDSLChannelPacketsReceived = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_dsl_packets_received_total",
		Help: "Number of packets received on the DSL channel",
	})

	// Wifi metrics
	ibWifiBytesSent = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_total_sent_bytes",
		Help: "Bytes sent on the WiFi channel, per device",
	}, []string{"channel", "station"})
	ibWifiBytesReceived = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_total_received_bytes",
		Help: "Bytes received on the WiFi channel per device",
	}, []string{"channel", "station"})
	ibWifiSignalStrength = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_signal_strength_db",
		Help: "Signal strength in dB per channel, per device",
	}, []string{"channel", "station"})
	ibWifiSignalNoiseRatio = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_signal_noise_ratio",
		Help: "Signal to noise ratio per channel, per device",
	}, []string{"channel", "station"})
	ibWifiBytesSentByBand = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_band_total_sent_bytes",
		Help: "Total sent bytes by WiFi band",
	}, []string{"band"})
	ibWifiBytesReceivedByBand = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_band_total_received_bytes",
		Help: "Total received bytes by WiFi band",
	}, []string{"band"})
	ibWifiDataMbps = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_data_mbps",
		Help: "Megabytes per second on data channel, by band and station",
	}, []string{"band", "station"})
	ibWifiPhyMbps = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_phy_mbps",
		Help: "Megabytes per second on physical channel, by band and station",
	}, []string{"band", "station"})
	ibWifiAirUse = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_air_use_percentage",
		Help: "Air use percentage, by band and station",
	}, []string{"band", "station"})
	ibWifiDataUse = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_wifi_data_use_percentage",
		Help: "Data use percentage, by band and station",
	}, []string{"band", "station"})

	// Connected devices
	ibClientsNumberOfActiveDevices = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "internetbox_client_devices_count",
		Help: "Currently connected devices, by interface",
	})
	ibClientRxBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_client_total_rx_bytes",
		Help: "Receiving bytes, by client",
	}, []string{"station", "ip", "interface"})
	ibClientTxBytes = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_client_total_tx_bytes",
		Help: "Transmitted bytes, by client",
	}, []string{"station", "ip", "interface"})
	ibClientSignalStrength = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_client_signalstrength",
		Help: "Signal strength",
	}, []string{"station", "ip", "interface"})
	ibClientNoise = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "internetbox_client_noise",
		Help: "Client noise",
	}, []string{"station", "ip", "interface"})
)
