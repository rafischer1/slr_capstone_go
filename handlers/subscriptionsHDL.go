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
	fmt.Println("Hit the getAll messages route:", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Write(resData)

}

// PostSub is a function
func PostSub(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("In the handler post req.Body: %+v", req.Method)
	if req.Method == "OPTIONS" {
		fmt.Println("Options in POST:", req.Method)
	}
	if req.Method == "POST" {
		enableCors(&w)
		fmt.Println("header in POST req:", &w)
		if req.Body != nil {
			bodyBytes, _ = ioutil.ReadAll(req.Body)
		}

		// Restore the io.ReadCloser to its original state
		req.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		// Use the content
		bodyString := string(bodyBytes)
		// body := m.Message{}
		str := bodyString
		res := m.Subscriber{}
		json.Unmarshal([]byte(str), &res)
		fmt.Println("res phone:", res.Phone, "res location:", res.Location)

		err := m.PostSub(res.Phone, res.Location)
		if err != nil {
			fmt.Fprint(w, "Content:", err, res.Phone, res.Location)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Content:", err, res.Phone, res.Location)
	}
}

// DeleteSub sends the delete request to the db
func DeleteSub(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	reqPhone := req.URL.String()
	phone := strings.Split(reqPhone, "/")[2]
	data, err := m.DeleteSub(phone)
	if err != nil {
		fmt.Fprint(w, "Error on delete:", err)
	}
	// vars := mux.Vars(req)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted Entry:", data)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS,DELETE,PUT")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
}
