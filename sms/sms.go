/**
* An Event happens from the front end there is a new message to be sent.
* It hits a route that triggers a database call to retrieve a list of numbers.
* The message content makes it to the request call in the correct format
* A for loop calls the sms function for every number in the db
* general hilarity ensues 🌊
 */

package sms

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	m "github.com/rafischer1/slr_capstone_go/models"
)

// if there is a POST to the /data route the handler could trigger first a GetAll call to the subscribers table and pass that []string of numbers as one of two parameters to the SendText(numbers, msg) => return twillio http.statust

// SendText formats and passes the flooding event message to the Twillio api
func SendText(Msg string) {
	accountSid, authToken, numberFrom := Init()

	subscriptions := m.GetAllSubs()
	for i, sub := range subscriptions {
		fmt.Printf("SMS to Number**%v = %v\n", i, sub.Phone)
	}

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// how to load all the data needed here?
	msgData := url.Values{}

	msgData.Set("To", "3024232120")
	msgData.Set("From", numberFrom)
	msgData.Set("Body", Msg)
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

// SubscribeSMS takes a phone number and send a subscription succesfull SMS
func SubscribeSMS(Phone string) {
	accountSid, authToken, numberFrom := Init()

	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	subMsg := Phone + "sucessfully subscribed to SLR Maine. TO unsubscribe at any time please visit the visit the slr-maine site. - - Global sea level has risen by about 8 inches since reliable record keeping began in 1880. Scientists predict sea level will rise between 8 inches and 6.6 feet by 2100."

	msgData := url.Values{}

	msgData.Set("To", "3024232120")
	msgData.Set("From", numberFrom)
	msgData.Set("Body", subMsg)
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

// Init initializes the SMS/Twillio env variables
func Init() (string, string, string) {
	accountSid := os.Getenv("SMS_SID")
	authToken := os.Getenv("SMS_TOKEN")
	numberFrom := os.Getenv("SMS_NUMBER")
	return accountSid, authToken, numberFrom
}
