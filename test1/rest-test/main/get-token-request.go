package main

import (
	"bytes"
	"log"
	"net/http"
)

func getTokenRequest() *http.Request {
	url := "http://backend-testing.emotorad.com/rest-test1/get-token" // "http://localhost:8081/query"
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(""))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req
}
