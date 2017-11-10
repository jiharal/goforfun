package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SendSmsUsingTwilio(pesan, numberTo string) {
	accountSid := "YOURACOUNT"
	authToken := "TOKEN"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"
	// Set data sms
	msgData := url.Values{}
	msgData.Set("To", numberTo)
	msgData.Set("From", "+19472227893")
	msgData.Set("Body", pesan)
	msgDataReader := *strings.NewReader(msgData.Encode())
	// Req sms
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}

func main() {
	numberTo := "+6282325600996"
	SendSmsUsingTwilio("Hello jihar", numberTo)
}
