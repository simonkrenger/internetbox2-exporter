package nmc

// LanIPResponse returned by /sysbus/NMC:getLANIP
type LanIPResponse struct {
	Status interface{} `json:"status"`
	Data   struct {
		Address        string `json:"Address"`
		Netmask        string `json:"Netmask"`
		DHCPEnable     bool   `json:"DHCPEnable"`
		DHCPMinAddress string `json:"DHCPMinAddress"`
		DHCPMaxAddress string `json:"DHCPMaxAddress"`
		LeaseTime      int    `json:"LeaseTime"`
	} `json:"data"`
}

// WanStatusResponse returned by /sysbus/NMC:getWANStatus
type WanStatusResponse struct {
	Status bool `json:"status"`
	Data   struct {
		LinkType        string `json:"LinkType"`
		LinkState       string `json:"LinkState"`
		Protocol        string `json:"Protocol"`
		ConnectionState string `json:"ConnectionState"`
	} `json:"data"`
}

// WifiDriverStatsResponse returned by /sysbus/NMC/Debug:getWifiDriverStats
type WifiDriverStatsResponse struct {
	Status struct {
		Band24G struct {
			Parameters struct {
				InterferenceStatus string `json:"InterferenceStatus"`
				Temperature        string `json:"Temperature"`
				ChipState          string `json:"ChipState"`
			} `json:"Parameters"`
			ChanimStats struct {
				Chanspec  string `json:"chanspec"`
				Tx        int    `json:"tx"`
				Inbss     int    `json:"inbss"`
				Obss      int    `json:"obss"`
				Nocat     int    `json:"nocat"`
				Nopkt     int    `json:"nopkt"`
				Doze      int    `json:"doze"`
				Txop      int    `json:"txop"`
				Goodtx    int    `json:"goodtx"`
				Badtx     int    `json:"badtx"`
				Glitch    int    `json:"glitch"`
				Badplcp   int    `json:"badplcp"`
				Knoise    int    `json:"knoise"`
				Idle      int    `json:"idle"`
				Timestamp int    `json:"timestamp"`
			} `json:"ChanimStats"`
			RssiPerAnt []struct {
				Value   string `json:"Value"`
				Average int    `json:"Average"`
			} `json:"RssiPerAnt"`
			LatestPower []struct {
				Value string `json:"Value"`
			} `json:"LatestPower"`
			Counters    WifiDriverStatsCounters   `json:"Counters"`
			Stations    []WifiDriverStatsStations `json:"Stations"`
			PacketQueue []struct {
				Values string `json:"Values"`
			} `json:"PacketQueue"`
			TemperatureThreshold struct {
				CurrentThreshold int    `json:"CurrentThreshold"`
				Hardware         string `json:"Hardware"`
				EqualUpper       int    `json:"EqualUpper"`
				Lower            int    `json:"Lower"`
			} `json:"TemperatureThreshold"`
		} `json:"Band24G"`
		Band5G struct {
			Parameters struct {
				InterferenceStatus string `json:"InterferenceStatus"`
				Temperature        string `json:"Temperature"`
				ChipState          string `json:"ChipState"`
			} `json:"Parameters"`
			ChanimStats struct {
				Chanspec  string `json:"chanspec"`
				Tx        int    `json:"tx"`
				Inbss     int    `json:"inbss"`
				Obss      int    `json:"obss"`
				Nocat     int    `json:"nocat"`
				Nopkt     int    `json:"nopkt"`
				Doze      int    `json:"doze"`
				Txop      int    `json:"txop"`
				Goodtx    int    `json:"goodtx"`
				Badtx     int    `json:"badtx"`
				Glitch    int    `json:"glitch"`
				Badplcp   int    `json:"badplcp"`
				Knoise    int    `json:"knoise"`
				Idle      int    `json:"idle"`
				Timestamp int    `json:"timestamp"`
			} `json:"ChanimStats"`
			RssiPerAnt []struct {
				Value   string `json:"Value"`
				Average int    `json:"Average"`
			} `json:"RssiPerAnt"`
			LatestPower []struct {
				Value string `json:"Value"`
			} `json:"LatestPower"`
			Counters    WifiDriverStatsCounters   `json:"Counters"`
			Stations    []WifiDriverStatsStations `json:"Stations"`
			PacketQueue []struct {
				Values string `json:"Values"`
			} `json:"PacketQueue"`
			TemperatureThreshold struct {
				CurrentThreshold int    `json:"CurrentThreshold"`
				Hardware         string `json:"Hardware"`
				EqualUpper       int    `json:"EqualUpper"`
				Lower            int    `json:"Lower"`
			} `json:"TemperatureThreshold"`
		} `json:"Band5G"`
	}
}

type WifiDriverStatsStations struct {
	Mac            string `json:"MAC"`
	ConnectedSince string `json:"ConnectedSince"`
	Flags          string `json:"flags"`
	PHYMbps        string `json:"PHYMbps"`
	DataMbps       string `json:"DataMbps"`
	AirUse         string `json:"AirUse"`
	DataUse        string `json:"DataUse"`
	Retries        string `json:"Retries"`
}

type WifiDriverStatsCounters struct {
	Txframe   int   `json:"txframe"`
	Txbyte    int64 `json:"txbyte"`
	Txretrans int   `json:"txretrans"`
	Txerror   int   `json:"txerror"`
	Rxframe   int   `json:"rxframe"`
	Rxbyte    int64 `json:"rxbyte"`
	Rxerror   int   `json:"rxerror"`
	Txprshort int   `json:"txprshort"`
	Txdmawar  int   `json:"txdmawar"`
	Txnobuf   int   `json:"txnobuf"`
	Txnoassoc int   `json:"txnoassoc"`
	Txchit    int   `json:"txchit"`
	Txcmiss   int   `json:"txcmiss"`
	Reset     int   `json:"reset"`
	Txserr    int   `json:"txserr"`
	Txphyerr  int   `json:"txphyerr"`
	Txphycrs  int   `json:"txphycrs"`
	Txlost    int   `json:"txlost"`
	Txfail    int   `json:"txfail"`
	Txchanrej int   `json:"txchanrej"`
	Tbtt      int   `json:"tbtt"`
}
