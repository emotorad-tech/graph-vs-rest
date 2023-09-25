package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Token struct {
	Token string `json:"token"`
}

var token Token

func getToken(req *http.Request) string {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&token)
	if err != nil {
		log.Panic(err)
	}

	return token.Token
}
