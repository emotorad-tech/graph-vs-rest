package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type TokenResponse struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
}

var tokenResp TokenResponse

func checkToken(req *http.Request) (string, bool) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&tokenResp)
	if err != nil {
		log.Panic(err)
	}

	return tokenResp.Token, tokenResp.Valid
}
