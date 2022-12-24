package sysbus

import "time"

// DeviceInfoResponse from /sysbus/DeviceInfo:get
type DeviceInfoResponse struct {
	Status struct {
		Manufacturer                             string    `json:"Manufacturer"`
		ManufacturerOUI                          string    `json:"ManufacturerOUI"`
		ModelName                                string    `json:"ModelName"`
		Description                              string    `json:"Description"`
		ProductClass                             string    `json:"ProductClass"`
		SerialNumber                             string    `json:"SerialNumber"`
		HardwareVersion                          string    `json:"HardwareVersion"`
		SoftwareVersion                          string    `json:"SoftwareVersion"`
		RescueVersion                            string    `json:"RescueVersion"`
		ModemFirmwareVersion                     string    `json:"ModemFirmwareVersion"`
		EnabledOptions                           string    `json:"EnabledOptions"`
		AdditionalHardwareVersion                string    `json:"AdditionalHardwareVersion"`
		AdditionalSoftwareVersion                string    `json:"AdditionalSoftwareVersion"`
		SpecVersion                              string    `json:"SpecVersion"`
		ProvisioningCode                         string    `json:"ProvisioningCode"`
		UpTime                                   int       `json:"UpTime"`
		FirstUseDate                             time.Time `json:"FirstUseDate"`
		DeviceLog                                string    `json:"DeviceLog"`
		VendorConfigFileNumberOfEntries          int       `json:"VendorConfigFileNumberOfEntries"`
		ManufacturerURL                          string    `json:"ManufacturerURL"`
		Country                                  string    `json:"Country"`
		ExternalIPAddress                        string    `json:"ExternalIPAddress"`
		DeviceStatus                             string    `json:"DeviceStatus"`
		NumberOfReboots                          int       `json:"NumberOfReboots"`
		UpgradeOccurred                          bool      `json:"UpgradeOccurred"`
		ResetOccurred                            bool      `json:"ResetOccurred"`
		RestoreOccurred                          bool      `json:"RestoreOccurred"`
		XSOFTATHOMECOMAdditionalSoftwareVersions string    `json:"X_SOFTATHOME-COM_AdditionalSoftwareVersions"`
		BaseMAC                                  string    `json:"BaseMAC"`
	} `json:"status"`
}

// GenLogsNextLogsResponse is returned by /sysbus/GenLog:nextLogs
type GenLogsNextLogsResponse struct {
	Status []struct {
		Timestamp time.Time `json:"timestamp"`
		Source    string    `json:"source"`
		Topic     string    `json:"topic"`
		Params    struct {
			TxBytes int64 `json:"TxBytes"`
			RxBytes int64 `json:"RxBytes"`
		} `json:"params"`
		Meta struct {
			Stream   string `json:"_stream"`
			Uniqueid int    `json:"_uniqueid"`
		} `json:"_meta"`
	} `json:"status"`
	Data struct {
		Iterator struct {
			Stream       string        `json:"_stream"`
			UniqueidList []interface{} `json:"_uniqueid_list"`
		} `json:"iterator"`
	} `json:"data"`
}

// Response from /sysbus/GenLog:readLogs
type GenLogReadLogResponse struct {
	Status struct {
		Stream       string `json:"_stream"`
		UniqueidList []int  `json:"_uniqueid_list"`
	} `json:"status"`
}

type NextLogRequest struct {
	Parameters NextLogRequestParameters `json:"parameters"`
}
type NextLogRequestParameters struct {
	Iterator NextLogRequestParametersIterator `json:"iterator"`
	Nrecords int                              `json:"nrecords"`
}
type NextLogRequestParametersIterator struct {
	Stream       string `json:"_stream"`
	UniqueidList []int  `json:"_uniqueid_list"`
}

type ReadLogRequest struct {
	ReadLogRequestParameters `json:"parameters"`
}

type ReadLogRequestParameters struct {
	Nrecords int    `json:"nrecords"`
	Since    string `json:"since"`
	Source   string `json:"source"`
	Topic    string `json:"topic"`
	Until    string `json:"until"`
}
