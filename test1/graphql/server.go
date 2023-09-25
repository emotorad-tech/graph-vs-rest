package main

import (
	"graphql/graph"
	"graphql/middlewares"
	"graphql/token"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
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

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middlewares.TokenMakerMiddleware(tokenMaker, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
