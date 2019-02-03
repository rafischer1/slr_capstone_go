package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	m "github.com/rafischer1/slr_capstone_go/models"
)

// Reader type interface
type Reader interface {
	Read(buf []byte) (n int, err error)
}

var bodyBytes []byte

// GetAllSubs handler to handle all records
func GetAllSubs(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	// fmt.Println("in the getall handler", req)
	data := m.GetAllSubs()

	//return the data
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)

}

// PostSub is a function
func PostSub(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	if req.Method == "OPTIONS" {
		fmt.Println("Options in POST:", req.Method)
	}
	if req.Method == "POST" {

		if req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(req.Body)
		}

		// Restore the io.ReadCloser to its original state
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)
		str := bodyString
		res := m.Subscriber{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println("res phone:", res.Phone, "res location:", res.Location)

		err := m.PostSub(res.Phone, res.Location)
		if err != nil {
			//send the error as JSON
			json.NewEncoder(w).Encode(err)
		} else {
			sms.SubscribeSMS(Phone)
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			// encode the response for JSON on the frontened
			json.NewEncoder(w).Encode(http.StatusOK)
		}

	}
}

// DeleteSub sends the delete request to the Delete model
func DeleteSub(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	reqPhone := req.URL.String()
	phone := strings.Split(reqPhone, "/")[2]
	_, err := m.DeleteSub(phone)
	// fmt.Printf("%T", res)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}

// MyError is used to Define Error struct
type MyError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *MyError) Error() string {
	return error.msg
}
