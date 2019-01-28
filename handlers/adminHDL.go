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

	// send password as param to the get admin SQL query result
	user, err := m.GetAdmin(password)

	//return the user
	w.WriteHeader(http.StatusOK)
	fmt.Println("Hit the admin verify route:", http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	// resUser, err := json.Marshal(user)

	// fmt.Println(len(user), "len user", resUser)
	if len(user) == 0 {
		json.NewEncoder(w).Encode(http.StatusBadRequest)
	} else if err != nil {
		json.NewEncoder(w).Encode(http.StatusNetworkAuthenticationRequired)
	} else {
		// w.Write(resUser)
		json.NewEncoder(w).Encode(http.StatusOK)
	}
}
