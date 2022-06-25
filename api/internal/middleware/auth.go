package middleware

import (
	"context"
	"ddgf-new/internal/repo"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the session ID cookie
		sid, err := r.Cookie("ddgf_sid")
		if err != nil {
			http.Error(w, "Forbidden. Please login to continue.", http.StatusForbidden)
			return
		}

		// Validate the sid
		_, err = uuid.Parse(sid.Value)
		if err != nil {
			http.Error(w, "Forbidden. Please login to continue", http.StatusForbidden)
			return
		}

		rc := r.Context().Value(RedisConn).(*redis.Client)

		// Get it from Redis if found
		val, err := rc.Get(context.Background(), sid.Value).Result()
		if err == redis.Nil {
			http.Error(w, "Forbidden. Please login to continue", http.StatusForbidden)
			return
		}

		// Decode the string from redis to session object
		var session repo.Session
		err = json.Unmarshal([]byte(val), &session)
		if err != nil {
			http.Error(w, "Error reading your session. Please reloging.", http.StatusInternalServerError)
			return
		}

		// If found, continue normal execution
		log.Printf("UserID: %v\n", session.UID)
		log.Printf("Role: %v\n", session.Role)
		next.ServeHTTP(w, r)
	})
}
