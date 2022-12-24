package ws

import "time"

// LoginRequest to log in
type LoginRequest struct {
	Service    string                 `json:"service"`
	Method     string                 `json:"method"`
	Parameters LoginRequestParameters `json:"parameters"`
}

// LoginRequestParameters used in LoginRequest
type LoginRequestParameters struct {
	ApplicationName string `json:"applicationName"`
	Username        string `json:"username"`
	Password        string `json:"password"`
}

// LoginResponse is the response when logging in
type LoginResponse struct {
	Status int                     `json:"status"`
	Data   LoginResponseParameters `json:"data"`
}

// LoginResponseParameters used in LoginResponse
type LoginResponseParameters struct {
	ContextID string `json:"contextID"`
	Username  string `json:"username"`
	Groups    string `json:"groups"`
}

type WSRequest struct {
	Service    string              `json:"service"`
	Method     string              `json:"method"`
	Parameters WSRequestParameters `json:"parameters"`
}

type WSRequestParameters struct {
	Channelid  int                 `json:"channelid,omitempty"`
	Expression string              `json:"expression,omitempty"`
	Flags      string              `json:"flags,omitempty"`
	Traverse   string              `json:"traverse,omitempty"`
	Type       string              `json:"type,omitempty"`
	Events     []WSEventsParameter `json:"events,omitempty"`
}

type WSGetDeviceResponse struct {
	Status []struct {
		Key                  string    `json:"Key"`
		DiscoverySource      string    `json:"DiscoverySource"`
		Name                 string    `json:"Name"`
		DeviceType           string    `json:"DeviceType"`
		Active               bool      `json:"Active"`
		Tags                 string    `json:"Tags"`
		FirstSeen            time.Time `json:"FirstSeen"`
		LastConnection       time.Time `json:"LastConnection"`
		LastChanged          time.Time `json:"LastChanged"`
		Master               string    `json:"Master"`
		VendorClassID        string    `json:"VendorClassID,omitempty"`
		UserClassID          string    `json:"UserClassID,omitempty"`
		ClientID             string    `json:"ClientID,omitempty"`
		SerialNumber         string    `json:"SerialNumber,omitempty"`
		ProductClass         string    `json:"ProductClass,omitempty"`
		Oui                  string    `json:"OUI,omitempty"`
		IPAddress            string    `json:"IPAddress,omitempty"`
		IPAddressSource      string    `json:"IPAddressSource,omitempty"`
		Location             string    `json:"Location,omitempty"`
		PhysAddress          string    `json:"PhysAddress,omitempty"`
		Layer2Interface      string    `json:"Layer2Interface,omitempty"`
		InterfaceName        string    `json:"InterfaceName,omitempty"`
		MACVendor            string    `json:"MACVendor,omitempty"`
		Owner                string    `json:"Owner,omitempty"`
		SignalStrength       int       `json:"SignalStrength,omitempty"`
		SignalNoiseRatio     int       `json:"SignalNoiseRatio,omitempty"`
		LastDataDownlinkRate int       `json:"LastDataDownlinkRate,omitempty"`
		LastDataUplinkRate   int       `json:"LastDataUplinkRate,omitempty"`
		EncryptionMode       string    `json:"EncryptionMode,omitempty"`
		LinkBandwidth        string    `json:"LinkBandwidth,omitempty"`
		SecurityModeEnabled  string    `json:"SecurityModeEnabled,omitempty"`
		HtCapabilities       string    `json:"HtCapabilities,omitempty"`
		VhtCapabilities      string    `json:"VhtCapabilities,omitempty"`
		HeCapabilities       string    `json:"HeCapabilities,omitempty"`
		SupportedMCS         string    `json:"SupportedMCS,omitempty"`
		Index                string    `json:"Index"`
		Actions              []struct {
			Function  string `json:"Function"`
			Name      string `json:"Name"`
			Arguments []struct {
				Name      string `json:"Name"`
				Type      string `json:"Type"`
				Mandatory bool   `json:"Mandatory"`
			} `json:"Arguments,omitempty"`
		} `json:"Actions,omitempty"`
		Names []struct {
			Name   string `json:"Name"`
			Source string `json:"Source"`
			Suffix string `json:"Suffix"`
			ID     string `json:"Id"`
		} `json:"Names"`
		DeviceTypes []struct {
			Type   string `json:"Type"`
			Source string `json:"Source"`
			ID     string `json:"Id"`
		} `json:"DeviceTypes"`
		IPv4Address []struct {
			Address       string `json:"Address"`
			Status        string `json:"Status"`
			Scope         string `json:"Scope"`
			AddressSource string `json:"AddressSource"`
			Reserved      bool   `json:"Reserved"`
			ID            string `json:"Id"`
		} `json:"IPv4Address,omitempty"`
		IPv6Address []struct {
			Address       string `json:"Address"`
			Status        string `json:"Status"`
			Scope         string `json:"Scope"`
			AddressSource string `json:"AddressSource"`
			ID            string `json:"Id"`
		} `json:"IPv6Address,omitempty"`
		Locations []interface{} `json:"Locations,omitempty"`
		Groups    []interface{} `json:"Groups,omitempty"`
		SSWSta    struct {
			SupportedStandards string `json:"SupportedStandards"`
			Supports24GHz      bool   `json:"Supports24GHz"`
			Supports5GHz       bool   `json:"Supports5GHz"`
			ReconnectClass     string `json:"ReconnectClass"`
			FailedSteerCount   int    `json:"FailedSteerCount"`
			SuccessSteerCount  int    `json:"SuccessSteerCount"`
			AvgSteeringTime    int    `json:"AvgSteeringTime"`
			SupportedUNIIBands string `json:"SupportedUNIIBands"`
		} `json:"SSWSta,omitempty"`
		MDNSService []struct {
			Name        string `json:"Name"`
			ServiceName string `json:"ServiceName"`
			Domain      string `json:"Domain"`
			Port        string `json:"Port"`
			ID          string `json:"Id"`
		} `json:"mDNSService,omitempty"`
		PortState              string `json:"PortState,omitempty"`
		NetDevName             string `json:"NetDevName,omitempty"`
		NetDevIndex            int    `json:"NetDevIndex,omitempty"`
		DHCPv4ServerPool       string `json:"DHCPv4ServerPool,omitempty"`
		DHCPv4ServerEnable     bool   `json:"DHCPv4ServerEnable,omitempty"`
		DHCPv4ServerMinAddress string `json:"DHCPv4ServerMinAddress,omitempty"`
		DHCPv4ServerMaxAddress string `json:"DHCPv4ServerMaxAddress,omitempty"`
		DHCPv4ServerNetmask    string `json:"DHCPv4ServerNetmask,omitempty"`
		DHCPv4DomainName       string `json:"DHCPv4DomainName,omitempty"`
		Ssid                   string `json:"SSID,omitempty"`
		Bssid                  string `json:"BSSID,omitempty"`
		OperatingFrequencyBand string `json:"OperatingFrequencyBand,omitempty"`
		OperatingStandards     string `json:"OperatingStandards,omitempty"`
		Channel                int    `json:"Channel,omitempty"`
		MaxBitRateSupported    int    `json:"MaxBitRateSupported,omitempty"`
		CurrentBitRate         int    `json:"CurrentBitRate,omitempty"`
	} `json:"status"`
}

type WSEventsParameter struct {
	Event   string `json:"event"`
	Service string `json:"service"`
	Data    struct {
		Types string `json:"types"`
	} `json:"data"`
}

type WSGetEventsRequest struct {
	Parameters WSGetEventsRequestParameters `json:"parameters"`
	Service    string                       `json:"service"`
	Method     string                       `json:"method"`
}

type WSGetEventsRequestParameters struct {
	Events    []WSGetEventsRequestEvent `json:"events"`
	Channelid int                       `json:"channelid"`
}

type WSGetEventsRequestEvent struct {
	Service string `json:"service"`
	Event   string `json:"event"`
	Data    struct {
		Types string `json:"types,omitempty"`
	} `json:"data,omitempty"`
}

type WSGetEventsResponse struct {
	Status struct {
		Channelid int `json:"channelid"`
		Events    []struct {
			Data struct {
				Handler string `json:"handler"`
				Object  struct {
					Reason     string `json:"reason"`
					Attributes struct {
						Hgw struct {
							Timestamp        time.Time `json:"Timestamp"`
							Uptime           int       `json:"uptime"`
							Loadaverage1Min  int       `json:"loadaverage_1min"`
							Loadaverage5Min  int       `json:"loadaverage_5min"`
							Loadaverage15Min int       `json:"loadaverage_15min"`
							Totalram         int       `json:"totalram"`
							Freeram          int       `json:"freeram"`
							Sharedram        int       `json:"sharedram"`
							Bufferram        int       `json:"bufferram"`
							Cachedram        int       `json:"cachedram"`
							Totalswap        int       `json:"totalswap"`
							Freeswap         int       `json:"freeswap"`
							Procs            int       `json:"procs"`
							Totalhigh        int       `json:"totalhigh"`
							MemUnit          int       `json:"mem_unit"`
						} `json:"HGW"`
					} `json:"attributes"`
				} `json:"object"`
			} `json:"data"`
		} `json:"events"`
	} `json:"status"`
}

type GetEventLogRequest struct {
	Method     string `json:"method"`
	Parameters struct {
		Type string `json:"type"`
	} `json:"parameters"`
	Service string `json:"service"`
}

type GetEventLogResponse struct {
	Status []struct {
		TxPackets        int       `json:"TxPackets"`
		RxPackets        int       `json:"RxPackets"`
		TxBytes          int       `json:"TxBytes"`
		RxBytes          int       `json:"RxBytes"`
		SignalStrength   int       `json:"SignalStrength"`
		SignalNoiseRatio int       `json:"SignalNoiseRatio"`
		Noise            int       `json:"Noise"`
		Timestamp        time.Time `json:"Timestamp"`
	} `json:"status"`
}
