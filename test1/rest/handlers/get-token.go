package handlers

import (
	"net/http"
	"rest/middlewares"
	"time"
)

type ResponsePayload struct {
	Token string `json:"token"`
}

func GetToken(w http.ResponseWriter, r *http.Request) {
	tokenMaker, err := middlewares.GetTokenMakerFromContext(r.Context())
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}

	token, err := tokenMaker.CreateToken("test", time.Minute*30)
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}

	writeJson(w, http.StatusOK, ResponsePayload{Token: token})
}
