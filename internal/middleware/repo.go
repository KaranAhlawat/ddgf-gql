package middleware

import (
	"context"
	"ddgf-new/internal/repo"
	"net/http"

	"gorm.io/gorm"
)

type RepoContext string

var repoContext RepoContext = "Repository"

func RepoMiddleware(db *gorm.DB) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			pgRepo := repo.NewPSQLRepository(db)
			ctx := context.WithValue(r.Context(), repoContext, pgRepo)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
