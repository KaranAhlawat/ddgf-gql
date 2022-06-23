package main

import (
	"ddgf-new/internal/server"
	"log"
	"net/http"
)

func main() {
	s := server.NewServer()

	s.SetupMiddleware()
	s.SetupRoutes()

	log.Println("Listening on http://localhost:8080/api/graphiql")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
