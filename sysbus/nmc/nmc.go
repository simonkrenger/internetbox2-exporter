package nmc

import (
	"encoding/json"
	"internetbox2-golang/ws"
	"log"
	"strconv"
	"strings"
)

func GetWANStatus() (WanStatusResponse, error) {

	wanStatusResp, err := ws.PostWithEmptyParameters("/sysbus/NMC:getWANStatus")
	if err != nil {
		log.Fatalln(err)
	}
	var wanStatus WanStatusResponse
	err = json.Unmarshal(wanStatusResp, &wanStatus)
	if err != nil {
		return WanStatusResponse{}, err
	}
	return wanStatus, nil
}

func GetLANIP() (LanIPResponse, error) {

	lanIPResp, err := ws.PostWithEmptyParameters("/sysbus/NMC:getLANIP")
	if err != nil {
		log.Fatalln(err)
	}
	var lanIP LanIPResponse
	err = json.Unmarshal(lanIPResp, &lanIP)
	if err != nil {
		return LanIPResponse{}, err
	}
	return lanIP, nil
}

func GetWifiDriverStats() (WifiDriverStatsResponse, error) {

	wifiDriverStatsResp, err := ws.PostWithEmptyParameters("/sysbus/NMC/Debug:getWifiDriverStats")
	if err != nil {
		log.Fatalln(err)
	}
	var driverStats WifiDriverStatsResponse
	err = json.Unmarshal(wifiDriverStatsResp, &driverStats)
	if err != nil {
		return WifiDriverStatsResponse{}, err
	}
	return driverStats, nil
}

func ParseMbps(input string) float64 {
	mbps, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0
	}
	return mbps
}

func ParseTemperature(input string) int {
	// Example input: "50 (0x32)"
	temp, err := strconv.Atoi(input[0:2])
	if err != nil {
		log.Fatalf("Error parsing temperature: " + err.Error())
		return 0
	}
	return temp
}

func ParsePercentage(input string) float64 {
	f, err := strconv.ParseFloat(strings.Replace(input, "%", "", 1), 64)
	if err != nil {
		log.Fatalf("Error parsing percentage: " + err.Error())
		return 0
	}
	return f
}
