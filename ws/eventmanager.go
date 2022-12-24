package ws

import (
	"encoding/json"
	"log"
)

// Device can be a MAC address (so the request becomes "Devices.Device.8C:85:90:2E:4D:39")
// Device can also be "HGW"
func GetEventLogForDevice(device string) (GetEventLogResponse, error) {

	payload, err := json.Marshal(WSRequest{
		Service: "Devices.Device." + device,
		Method:  "getEventLog",
		Parameters: WSRequestParameters{
			Type: "network",
		},
	})
	if err != nil {
		log.Println(err)
	}
	logsResp, err := PostWithObject("/ws", payload)
	if err != nil {
		log.Println(err)
		return GetEventLogResponse{}, nil
	}
	var logs GetEventLogResponse
	err = json.Unmarshal(logsResp, &logs)
	if err != nil {
		log.Println(err)
		return GetEventLogResponse{}, nil
	}

	return logs, nil
}

func GetInternetBoxStatus() (WSGetEventsResponse, error) {

	payload, err := json.Marshal(WSGetEventsRequest{
		Service: "eventmanager",
		Method:  "get_events",
		Parameters: WSGetEventsRequestParameters{
			Channelid: 0,
			Events: []WSGetEventsRequestEvent{
				{Event: "cpu", Service: "Devices.Device.HGW"},
				{Event: "sysinfo", Service: "Devices.Device.HGW", Data: struct {
					Types string "json:\"types,omitempty\""
				}{"sysinfo"}}},
		},
	})
	if err != nil {
		log.Println(err)
	}
	eventsResp, err := PostWithObject("/ws", payload)
	if err != nil {
		log.Println(err)
		return WSGetEventsResponse{}, nil
	}
	var events WSGetEventsResponse
	err = json.Unmarshal(eventsResp, &events)
	if err != nil {
		log.Println(err)
		return WSGetEventsResponse{}, nil
	}

	return events, nil
}
