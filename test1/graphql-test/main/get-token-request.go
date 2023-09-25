package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func getTokenRequest() *http.Request {
	url := "http://backend-testing.emotorad.com/graphql-test1/query" // "http://localhost:8081/query"
	query := `
		query {
			getToken
		}
	`
	requestBody := map[string]interface{}{
		"query": query,
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
