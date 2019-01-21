package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	m "github.com/rafischer1/slr_capstone_go/models"
	"github.com/rafischer1/slr_capstone_go/sms"
)

// GetAllSubs handler to handle all records
func GetAllData(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	// fmt.Println("in the getall handler", req)
	data := m.GetAllData()

	//return the data
	w.WriteHeader(http.StatusOK)
	fmt.Println("Hit the getAll messages route:", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)

}

// PostData recieves data from a POST to /data w/ SMS message info
func PostData(w http.ResponseWriter, req *http.Request) {
	if req.Method == "OPTIONS" {
		fmt.Println("Options in data POST:", req.Method)
	}
	if req.Method == "POST" {
		fmt.Println("header in data POST req:", &w)

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
		fmt.Println("the whole dang res:", res)

		fmt.Println("res data:", res.Msg, "res windmph:", res.WindMPH, "res winddir:", res.WindDir, "Sea level ft:", res.SeaLevelFt)

		defer sms.SendText(res.Msg)
		err := m.PostData(res.Msg, res.WindMPH, res.WindDir, res.SeaLevelFt)
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
