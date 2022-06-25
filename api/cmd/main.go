package main

import (
	"ddgf-new/internal/repo"
	"ddgf-new/internal/server"
	"log"
	"net/http"
)

func main() {
	db := repo.InitPostgresConn()
	pr := repo.NewPSQLRepository(db)
	rr := repo.NewRedisClient()

	s := server.NewServer()

	s.SetupRoutes(&pr, rr)

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
