package handlers

import (
	"net/http"
	"rest/middlewares"
)

type requestBody struct {
	Token string `json:"token"`
}

type responseBody struct {
	Token string `json:"token"`
	Valid bool   `json:"valid"`
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	var request requestBody
	readJsonFromBody(w, r, &request)

	tokenMaker, err := middlewares.GetTokenMakerFromContext(r.Context())
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}

	payload, err := tokenMaker.VerifyToken(request.Token)
	if err != nil {
		errorJson(w, err, http.StatusInternalServerError)
		return
	}

	if payload.Name != "test" {
		writeJson(w, http.StatusOK, responseBody{
			Token: request.Token,
			Valid: false,
		})
		return
	}

	writeJson(w, http.StatusOK, responseBody{Token: request.Token, Valid: true})
}
