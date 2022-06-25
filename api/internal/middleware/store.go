package middleware

import (
	"context"
	"ddgf-new/internal/repo"
	"net/http"

	"github.com/go-redis/redis/v9"
)

const RedisConn repo.Redis = "redis"

func RedisMiddleware(rc *redis.Client) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			ctx := context.WithValue(r.Context(), RedisConn, rc)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
