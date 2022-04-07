package watttime

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	endpoint = "https://api2.watttime.org"
	client   = &http.Client{}
)

func Login(username, password string) Token {
	url := endpoint + "/login"
	req, _ := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, password)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var token Token
	if err := json.Unmarshal(body, &token); err != nil {
		fmt.Println("Cannot unmarshall JSON")
	}
	return token
}

func Index(token Token, ba string) RealTimeEmissionsIndex {
	url := endpoint + "/index?ba=" + ba
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Bearer " + token.Value)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	var rtei RealTimeEmissionsIndex
	if err := json.Unmarshal(body, &rtei); err != nil {
		fmt.Println("Cannot unmarshall JSON")
	}
	return rtei
}