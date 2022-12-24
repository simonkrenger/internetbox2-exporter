package nemo

import (
	"encoding/json"
	"internetbox2-golang/ws"
	"log"
)

// Read statistics from the DSL Channel
func GetDSLChannelStats(channel string) (DSLChannelStatsResponse, error) {

	dslChannelStatsResp, err := ws.PostWithEmptyParameters("/sysbus/NeMo/Intf/" + channel + ":getDSLChannelStats")
	if err != nil {
		log.Fatalln(err)
	}
	var DSLChannelStats DSLChannelStatsResponse
	err = json.Unmarshal(dslChannelStatsResp, &DSLChannelStats)
	if err != nil {
		log.Fatalln(err)
		return DSLChannelStatsResponse{}, err
	}
	return DSLChannelStats, nil
}

// Retrieve WiFi statistics for a certain WiFi-connected station
func GetWifiStationStats(device string) (WifiChannelStatsResponse, error) {

	wifiChannelStatsResp, err := ws.PostWithEmptyParameters("/sysbus/NeMo/Intf/" + device + ":getStationStats")
	if err != nil {
		log.Fatalln(err)
	}
	var WifiChannelStats WifiChannelStatsResponse
	err = json.Unmarshal(wifiChannelStatsResp, &WifiChannelStats)
	if err != nil {
		log.Fatalln(err)
		return WifiChannelStatsResponse{}, err
	}
	return WifiChannelStats, nil
}
