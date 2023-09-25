package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest/middlewares"
	"rest/token"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	tokenMaker, err := token.NewPasetoMaker("1234567890123456789012345678901212345678901234567890123456789012")
	if err != nil {
		log.Panic("error creating token maker")
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: middlewares.TokenMakerMiddleware(tokenMaker, routes()),
	}

	log.Println("starting server on port", port)
	if srv.ListenAndServe() != nil {
		log.Panic("error starting server")
	}
}
