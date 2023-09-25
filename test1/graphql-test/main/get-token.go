package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func getToken(req *http.Request) string {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Panic(err)
	}

	dataField, ok := response["data"]
	if !ok {
		log.Panic("No data field in response")
	}

	getTokenField, ok := dataField.(map[string]interface{})["getToken"]
	if !ok {
		fmt.Println("Response 'data' field does not contain 'getToken'")
	}

	token, ok := getTokenField.(string)
	if !ok {
		fmt.Println("Failed to extract token as a string")
	}
	return token
}
