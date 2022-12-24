package sysbus

import (
	"encoding/json"
	"internetbox2-golang/ws"
	"log"
)

func GetDeviceInfo() (DeviceInfoResponse, error) {
	deviceInfoResp, err := ws.PostWithEmptyParameters("/sysbus/DeviceInfo:get")
	if err != nil {
		log.Fatalln(err)
	}
	var deviceInfo DeviceInfoResponse
	err = json.Unmarshal(deviceInfoResp, &deviceInfo)
	if err != nil {
		return DeviceInfoResponse{}, err
	}
	return deviceInfo, nil
}
