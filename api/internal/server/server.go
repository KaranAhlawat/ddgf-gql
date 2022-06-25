package server

import (
	"ddgf-new/internal/graph"
	"ddgf-new/internal/middleware"
	"ddgf-new/internal/repo"
	"ddgf-new/internal/resolver"
	"net/http"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-redis/redis/v9"
	"github.com/gorilla/mux"
)

type Server struct {
	Router *mux.Router
}

func NewServer() Server {
	r := mux.NewRouter()
	return Server{Router: r}
}

func (s *Server) SetupRoutes(db *repo.PSQLRepository, rc *redis.Client) {
	schema := graph.NewExecutableSchema(graph.Config{
		Resolvers: &resolver.Resolver{
			DB: db,
		},
	})
	sh := handler.NewDefaultServer(schema)

	s.Router.Use(middleware.LoggingMiddleware)
	s.Router.Use(middleware.RedisMiddleware(rc))
	s.Router.Handle("/login", http.HandlerFunc(LoginHandler)).Methods("GET", "POST")

	apiRouter := s.Router.PathPrefix("/api").Subrouter()

	apiRouter.Use(middleware.AuthMiddleware)
	apiRouter.Handle("/graphiql", playground.Handler("DDGF", "/api")).Methods("GET")
	apiRouter.Handle("/", sh).Methods("POST", "OPTIONS", "PUT", "DELETE")
}
