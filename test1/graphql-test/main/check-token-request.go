package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func checkTokenRequest(token string) *http.Request {
	url := "http://backend-testing.emotorad.com/graphql-test1/query" // "http://localhost:8081/query"
	mutation := `
		mutation {
			checkToken(input: "` + token + `") {
				token
				valid
			}
		}
	`
	requestBody := map[string]interface{}{
		"query": mutation,
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
