package main

import (
	"ddgf-new/internal/repo"
	"ddgf-new/internal/server"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func main() {
	db := repo.InitPostgresConn()
	r := repo.NewPSQLRepository(db)

	s := server.NewServer()

	sessions := make(map[uuid.UUID]repo.Session, 0)

	s.SetupRoutes(&r, &sessions)

	log.Println("Listening on http://localhost:8080/api/graphiql")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
