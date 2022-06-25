package server

import (
	"ddgf-new/internal/graph"
	"ddgf-new/internal/middleware"
	"ddgf-new/internal/repo"
	"ddgf-new/internal/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() Server {
	r := mux.NewRouter()
	return Server{Router: r}
}

func (s *Server) SetupRoutes(db *repo.PSQLRepository, sessions *map[uuid.UUID]repo.Session) {
	schema := graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{
			DB: db,
		},
	})
	sh := handler.NewDefaultServer(schema)

	loginHandler := LoginHandler(sessions)

	s.Router.Use(middleware.LoggingMiddleware)
	s.Router.Handle("/login", loginHandler).Methods("GET", "POST")

	apiRouter := s.Router.PathPrefix("/api").Subrouter()
	apiRouter.Use(middleware.AuthMiddleware(sessions))
	apiRouter.Handle("/graphiql", playground.Handler("DDGF", "/api")).Methods("GET")
	apiRouter.Handle("/", sh).Methods("POST", "OPTIONS", "PUT", "DELETE")
}
