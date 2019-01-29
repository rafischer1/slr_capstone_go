package main_test

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttp(t *testing.T) {
	//
	handler := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "[{ \"location\": \"Portland\"}]")
	}

	req := httptest.NewRequest("GET", "/subscribe", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.StatusCode)
	fmt.Println(resp.Header.Get("Content-Type"))
	fmt.Println(string(body))
}
