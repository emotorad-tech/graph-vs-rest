package main

import (
	"bytes"
	"log"
	"net/http"
)

type RequestBody struct {
	Token string `json:"token"`
}

func getDataRequest() *http.Request {
	url := "http://backend-testing.emotorad.com/rest-test2/get-data" // "http://localhost:8080/get-data"
	req, err := http.NewRequest("GET", url, bytes.NewBufferString(""))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req
}
