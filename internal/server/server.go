package server

import (
	"ddgf-new/internal/graph"
	"ddgf-new/internal/middleware"
	"ddgf-new/internal/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	Router *mux.Router
}

func NewServer() Server {
	r := mux.NewRouter()
	return Server{Router: r}
}

func (s *Server) SetupMiddleware(db *gorm.DB) {
	repoMiddleware := middleware.RepoMiddleware(db)
	s.Router.Use(middleware.LoggingMiddleware)
	s.Router.Use(repoMiddleware)
}

func (s *Server) SetupRoutes() {
	schema := graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{},
	})
	sh := handler.NewDefaultServer(schema)

	s.Router.Handle("/api/graphiql", playground.Handler("DDGF", "/api")).Methods("GET")
	s.Router.Handle("/api", sh).Methods("POST", "OPTIONS", "PUT", "DELETE")
}
