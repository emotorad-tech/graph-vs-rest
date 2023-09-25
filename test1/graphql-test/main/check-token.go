package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func checkToken(req *http.Request) (string, bool) {
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

	field, ok := dataField.(map[string]interface{})["checkToken"]
	if !ok {
		fmt.Println("Response 'data' field does not contain 'getToken'")
	}

	token, ok := field.(map[string]interface{})["token"]
	if !ok {
		log.Panic("Response 'checkToken' field does not contain 'token'")
	}

	valid, ok := field.(map[string]interface{})["valid"]
	if !ok {
		log.Panic("Response 'checkToken' field does not contain 'valid'")
	}

	return token.(string), valid.(bool)
}
