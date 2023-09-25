package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type RequestBody struct {
	Token string `json:"token"`
}

func checkTokenRequest(token string) *http.Request {
	url := "http://backend-testing.emotorad.com/rest-test1/check-token" // "http://localhost:8080/check-token"
	requestBody := RequestBody{
		Token: token,
	}
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		log.Panic(err)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Panic(err)
	}
	req.Header.Set("Content-Type", "application/json")

	return req
}
