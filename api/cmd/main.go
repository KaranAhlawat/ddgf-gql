package main

import (
	"ddgf-new/internal/repo"
	"ddgf-new/internal/server"
	"log"
	"net/http"
)

func main() {
	db := repo.InitPostgresConn()
	repo := repo.NewPSQLRepository(db)

	s := server.NewServer()

	s.SetupMiddleware()
	s.SetupRoutes(&repo)

	log.Println("Listening on http://localhost:8080/api/graphiql")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
