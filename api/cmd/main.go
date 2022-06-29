package main

import (
	"ddgf-new/config"
	"ddgf-new/internal/repo"
	"ddgf-new/internal/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("Unable to read config from .env file: %s\n", err.Error())
	}

	db := repo.InitPostgresConn(cfg)
	rr := repo.NewRedisClient(cfg)
	pr := repo.NewPSQLRepository(db)

	s := server.NewServer()

	s.SetupRoutes(&pr, rr)

	log.Println("Listening on http://localhost:8080/")
	log.Fatal(http.ListenAndServe(":8080", s.Router))
}
