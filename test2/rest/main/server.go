package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"rest/middlewares"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	session := PostgresConnect()

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: middlewares.DatabaseMiddleware(session, routes()),
	}

	log.Println("starting server on port", port)
	if srv.ListenAndServe() != nil {
		log.Panic("error starting server")
	}
}
