package main

import (
	"log"
	"net/http"
)

func getData(req *http.Request) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()
}
