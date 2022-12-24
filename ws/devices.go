package ws

import (
	"encoding/json"
	"log"
)

// Expression can be used to limit the result set
// Sample expressions are:
// - "lan"
// - "not interface"
// - "lan and not self and not interface"
func GetDevices(expression string) (WSGetDeviceResponse, error) {

	payload, err := json.Marshal(WSRequest{
		Service: "Devices",
		Method:  "get",
		Parameters: WSRequestParameters{
			Expression: expression,
		},
	})
	if err != nil {
		log.Println(err)
	}
	deviceResp, err := PostWithObject("/ws", payload)
	if err != nil {
		log.Fatalln(err)
	}
	var devices WSGetDeviceResponse
	err = json.Unmarshal(deviceResp, &devices)
	if err != nil {
		log.Fatalln(err)
		return WSGetDeviceResponse{}, nil
	}

	return devices, nil
}
