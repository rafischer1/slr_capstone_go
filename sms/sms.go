package sms

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
)

/**
* An Event happens from the front end there is a new message to be sent.
* It hits a route that triggers a database call to retrieve a list of numbers.
* The message content makes it to the request call in the correct format
* A for loop calls the sms function for every number in the db
* general hilarity ensues 🌊
 */

// SendText will hopefully work someday
func SendText() {
	accountSid, authToken, numberFrom := Init()

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// how to load all the data needed here?
	message := [1]string{"Hi there"}

	msgData := url.Values{}

	msgData.Set("To", "+3024232120")
	msgData.Set("From", numberFrom)
	msgData.Set("Body", message[rand.Intn(len(message))])
	msgDataReader := *strings.NewReader(msgData.Encode())

	fmt.Println("msgDataReader:", msgDataReader)

	// request body is formed and sent
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// response is read and evaluated
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println("response println:", resp, resp.Status)
	}
}

// Init initializes the SMS env variables
func Init() (string, string, string) {
	accountSid := os.Getenv("SMS_SID")
	authToken := os.Getenv("SMS_TOKEN")
	numberFrom := os.Getenv("SMS_NUMBER")
	return accountSid, authToken, numberFrom
}
