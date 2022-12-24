package ws

import "testing"

func TestGetDevices(t *testing.T) {
	devices, err := GetDevices("lan")
	if err != nil {
		t.Fail()
	}
	if len(devices.Status) < 10 {
		t.Fail()
	}
}
