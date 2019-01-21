package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	m "github.com/rafischer1/slr_capstone_go/models"
)

// AdminVerify handler to varify password on request
func AdminVerify(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
	reqPass := req.URL.String()
	password := strings.Split(reqPass, "/")[2]
	user := m.GetAdmin(password)
	json.Marshal(user)

	//return the user
	w.WriteHeader(http.StatusOK)
	fmt.Println("Hit the getAll messages route:", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	resUser, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}
	w.Write(resUser)
}
