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

	subMsg := Phone + " subscribed to SLR Maine. To unsubscribe please visit the visit the slr-maine site. - - 'Because of global warming that has already occurred and warming that is yet to occur due to the uncertain level of future emissions. Sea level rise is a certain impact of climate change; the questions are when, and how much, rather than if' NOAA - 2017"

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
