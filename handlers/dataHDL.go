package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	m "github.com/rafischer1/slr_capstone_go/models"
	"github.com/rafischer1/slr_capstone_go/sms"
)

// CRUD operations for data

// GetAllData handler to handle all recorded flooding events
func GetAllData(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	// retrieve data from the get all model
	data := m.GetAllData()

	//return the data
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)
}

// PostData recieves data from a POST to /data w/ SMS message info
func PostData(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	if req.Method == "OPTIONS" {
		fmt.Println("Options in data POST:", req.Method)
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
		res := m.Datum{}
		json.Unmarshal([]byte(str), &res)

		fmt.Println("res data:", res.Msg, "res windmph:", res.WindMPH, "res winddir:", res.WindDir, "Sea level ft:", res.SeaLevelFt, "Category:", res.Category)

		// send the message on to the Send Text handler for processing
		defer sms.SendText(res.Msg)

		// post the flooding event data to the database
		err := m.PostData(res.Msg, res.WindMPH, res.WindDir, res.SeaLevelFt, res.Category)
		if err != nil {
			//send the error as JSON
			json.NewEncoder(w).Encode(err)
		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			// encode the response for JSON on the frontened
			json.NewEncoder(w).Encode(err)
		}
	}
}

// DeleteData sends the delete request to the Delete model
func DeleteData(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	fmt.Println("Delete data hit", req.Method)
	reqID := req.URL.String()
	id := strings.Split(reqID, "/")[2]
	t, _ := strconv.Atoi(id)
	_, err := m.DeleteDatum(t)
	fmt.Printf("%T", id)
	if err != nil {
		json.NewEncoder(w).Encode(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}
