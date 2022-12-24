package nemo

import "time"

// DSLChannelStatsResponse returned by /sysbus/NeMo/Intf/dsl0:getDSLChannelStats
type DSLChannelStatsResponse struct {
	Status struct {
		Stats struct {
			BytesSent       int `json:"BytesSent"`
			BytesReceived   int `json:"BytesReceived"`
			PacketsSent     int `json:"PacketsSent"`
			PacketsReceived int `json:"PacketsReceived"`
			ErrorsSent      int `json:"ErrorsSent"`
			ErrorsReceived  int `json:"ErrorsReceived"`
			TotalStart      int `json:"TotalStart"`
			ShowtimeStart   int `json:"ShowtimeStart"`
		} `json:"stats"`
		Total        DSLChannelStatsResponseError `json:"Total"`
		Showtime     DSLChannelStatsResponseError `json:"Showtime"`
		LastShowtime DSLChannelStatsResponseError `json:"LastShowtime"`
		CurrentDay   DSLChannelStatsResponseError `json:"CurrentDay"`
		QuarterHour  DSLChannelStatsResponseError `json:"QuarterHour"`
	} `json:"status"`
}

type DSLChannelStatsResponseError struct {
	XTURFECErrors int `json:"XTURFECErrors"`
	XTUCFECErrors int `json:"XTUCFECErrors"`
	XTURHECErrors int `json:"XTURHECErrors"`
	XTUCHECErrors int `json:"XTUCHECErrors"`
	XTURCRCErrors int `json:"XTURCRCErrors"`
	XTUCCRCErrors int `json:"XTUCCRCErrors"`
}

// WifiChannelStatsResponse returned by /sysbus/NeMo/Intf/wl1:getStationStats
type WifiChannelStatsResponse struct {
	Status []struct {
		ChargeableUserID             string    `json:"ChargeableUserId"`
		MACAddress                   string    `json:"MACAddress"`
		AuthenticationState          bool      `json:"AuthenticationState"`
		LastDataDownlinkRate         int       `json:"LastDataDownlinkRate"`
		LastDataUplinkRate           int       `json:"LastDataUplinkRate"`
		SignalStrength               int       `json:"SignalStrength"`
		SignalStrengthHistory        string    `json:"SignalStrengthHistory"`
		SignalStrengthByChain        string    `json:"SignalStrengthByChain"`
		AvgSignalStrengthByChain     int       `json:"AvgSignalStrengthByChain"`
		AvgSignalStrength            int       `json:"AvgSignalStrength"`
		Retransmissions              int       `json:"Retransmissions"`
		Active                       bool      `json:"Active"`
		SignalNoiseRatio             int       `json:"SignalNoiseRatio"`
		Noise                        int       `json:"Noise"`
		Inactive                     int       `json:"Inactive"`
		RxPacketCount                int       `json:"RxPacketCount"`
		TxPacketCount                int       `json:"TxPacketCount"`
		RxUnicastPacketCount         int       `json:"RxUnicastPacketCount"`
		TxUnicastPacketCount         int       `json:"TxUnicastPacketCount"`
		RxMulticastPacketCount       int       `json:"RxMulticastPacketCount"`
		TxMulticastPacketCount       int       `json:"TxMulticastPacketCount"`
		TxBytes                      int       `json:"TxBytes"`
		RxBytes                      int       `json:"RxBytes"`
		TxErrors                     int       `json:"TxErrors"`
		RxRetransmissions            int       `json:"Rx_Retransmissions"`
		TxRetransmissions            int       `json:"Tx_Retransmissions"`
		TxRetransmissionsFailed      int       `json:"Tx_RetransmissionsFailed"`
		UplinkMCS                    int       `json:"UplinkMCS"`
		UplinkBandwidth              int       `json:"UplinkBandwidth"`
		UplinkShortGuard             bool      `json:"UplinkShortGuard"`
		DownlinkMCS                  int       `json:"DownlinkMCS"`
		DownlinkBandwidth            int       `json:"DownlinkBandwidth"`
		DownlinkShortGuard           bool      `json:"DownlinkShortGuard"`
		MaxRxSpatialStreamsSupported int       `json:"MaxRxSpatialStreamsSupported"`
		MaxTxSpatialStreamsSupported int       `json:"MaxTxSpatialStreamsSupported"`
		MaxDownlinkRateSupported     int       `json:"MaxDownlinkRateSupported"`
		MaxDownlinkRateReached       int       `json:"MaxDownlinkRateReached"`
		MaxUplinkRateSupported       int       `json:"MaxUplinkRateSupported"`
		MaxUplinkRateReached         int       `json:"MaxUplinkRateReached"`
		MaxBandwidthSupported        string    `json:"MaxBandwidthSupported"`
		Mode                         string    `json:"Mode"`
		OperatingStandard            string    `json:"OperatingStandard"`
		PowerSave                    bool      `json:"PowerSave"`
		Capabilities                 string    `json:"Capabilities"`
		ConnectionDuration           int       `json:"ConnectionDuration"`
		DeviceType                   string    `json:"DeviceType"`
		DevicePriority               int       `json:"DevicePriority"`
		LastStateChange              time.Time `json:"LastStateChange"`
		AssociationTime              time.Time `json:"AssociationTime"`
		MUGroupID                    int       `json:"MUGroupId"`
		MUUserPositionID             int       `json:"MUUserPositionId"`
		MUMimoTxPktsCount            int       `json:"MUMimoTxPktsCount"`
		MUMimoTxPktsPercentage       int       `json:"MUMimoTxPktsPercentage"`
		SupportedMCS                 string    `json:"SupportedMCS"`
		SupportedVhtMCS              string    `json:"SupportedVhtMCS"`
		SupportedHeMCS               string    `json:"SupportedHeMCS"`
		SupportedHe160MCS            string    `json:"SupportedHe160MCS"`
		VendorOUI                    string    `json:"VendorOUI"`
		SecurityModeEnabled          string    `json:"SecurityModeEnabled"`
		LinkBandwidth                string    `json:"LinkBandwidth"`
		EncryptionMode               string    `json:"EncryptionMode"`
		HtCapabilities               string    `json:"HtCapabilities"`
		VhtCapabilities              string    `json:"VhtCapabilities"`
		HeCapabilities               string    `json:"HeCapabilities"`
		FrequencyCapabilities        string    `json:"FrequencyCapabilities"`
		UNIIBandsCapabilities        string    `json:"UNIIBandsCapabilities"`
		ProbeReqCaps                 struct {
			SupportedMCS          string `json:"SupportedMCS"`
			VendorOUI             string `json:"VendorOUI"`
			SecurityModeEnabled   string `json:"SecurityModeEnabled"`
			LinkBandwidth         string `json:"LinkBandwidth"`
			EncryptionMode        string `json:"EncryptionMode"`
			HtCapabilities        string `json:"HtCapabilities"`
			VhtCapabilities       string `json:"VhtCapabilities"`
			HeCapabilities        string `json:"HeCapabilities"`
			FrequencyCapabilities string `json:"FrequencyCapabilities"`
		} `json:"ProbeReqCaps"`
	} `json:"status"`
}
