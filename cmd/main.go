package main

import (
	"ddgf-new/internal/graph"
	"ddgf-new/internal/resolver"

	"log"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

func main() {
	port := "8080"

	server := handler.NewDefaultServer(
		graph.NewExecutableSchema(graph.Config{Resolvers: &resolver.Resolver{}}),
	)

	http.Handle("/", playground.Handler("GraphQL Playground", "/api"))
	http.Handle("/api", server)

	log.Printf("Connect to http://localhost:%s for GraphiQL", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
