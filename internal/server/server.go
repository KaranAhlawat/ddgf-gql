package server

import (
	"ddgf-new/internal/graph"
	"ddgf-new/internal/middleware"
	"ddgf-new/internal/repo"
	"ddgf-new/internal/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() Server {
	r := mux.NewRouter()
	return Server{Router: r}
}

func (s *Server) SetupMiddleware() {
	s.Router.Use(middleware.LoggingMiddleware)
}

func (s *Server) SetupRoutes(db *repo.PSQLRepository) {
	schema := graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{
			DB: db,
		},
	})
	sh := handler.NewDefaultServer(schema)

	s.Router.Handle("/api/graphiql", playground.Handler("DDGF", "/api")).Methods("GET")
	s.Router.Handle("/api", sh).Methods("POST", "OPTIONS", "PUT", "DELETE")
}
