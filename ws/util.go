package ws

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

var internetBoxBaseURL string
var internetBoxAdminPassword string
var currentLoginContext LoginContext

type LoginContext struct {
	contextID    string
	cookieString string
}

func getLoginContext() LoginContext {

	// Happy path if login has already happened
	if currentLoginContext.contextID != "" {
		return currentLoginContext
	}

	log.Println("No context, logging in...")

	internetBoxAdminPassword = os.Getenv("INTERNETBOX_ADMIN_PASSWORD")
	if internetBoxAdminPassword == "" {
		log.Fatalln("No INTERNETBOX_ADMIN_PASSWORD set")
	}

	internetBoxBaseURL = os.Getenv("INTERNETBOX_BASEURL")
	if internetBoxBaseURL == "" {
		internetBoxBaseURL = "http://internetbox.home"
		log.Println("No INTERNETBOX_BASEURL set, falling back to '" + internetBoxBaseURL + "'")
	}

	loginJSON, err := json.Marshal(LoginRequest{
		Service: "sah.Device.Information",
		Method:  "createContext",
		Parameters: LoginRequestParameters{
			ApplicationName: "webui",
			Username:        "admin",
			Password:        internetBoxAdminPassword,
		},
	})
	if err != nil {
		log.Fatalln(err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", internetBoxBaseURL+"/ws", bytes.NewBuffer(loginJSON))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("Authorization", "X-Sah-Login")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var loginResp LoginResponse
	err = json.Unmarshal(responseData, &loginResp)
	if err != nil {
		log.Fatal(err)
	}

	currentLoginContext.cookieString = resp.Header.Get("Set-Cookie")
	currentLoginContext.contextID = loginResp.Data.ContextID

	return currentLoginContext
}

func PostWithEmptyParameters(path string) ([]byte, error) {
	return PostWithObject(path, []byte(`{"parameters":{}}`))
}

func PostWithObject(path string, payload []byte) ([]byte, error) {
	context := getLoginContext()

	log.Println("POST '" + path + "' with object: " + string(payload))

	url := internetBoxBaseURL + path

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		log.Println(err)
		return []byte{}, nil
	}
	req.Header.Add("Content-Type", "application/x-sah-ws-4-call+json")
	req.Header.Add("X-Context", context.contextID)
	req.Header.Add("Authorization", "X-Sah "+context.contextID)
	req.Header.Add("Cookie", context.cookieString)

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return []byte{}, nil
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte{}, nil
	}

	log.Println("Response: " + string(responseData))

	if resp.StatusCode != 200 {
		log.Println(err)
		return []byte{}, nil
	}

	return responseData, nil
}
