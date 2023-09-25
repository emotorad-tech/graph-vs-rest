package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type DataResponse struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
	Age   int    `json:"age"`
}

var dataResp []DataResponse

func getData(req *http.Request) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&dataResp)
	if err != nil {
		log.Panic(err)
	}
}
