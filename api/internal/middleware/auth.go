package middleware

import (
	"ddgf-new/internal/repo"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

func AuthMiddleware(sessions *map[uuid.UUID]repo.Session) func(next http.Handler) http.Handler {
	for k, v := range *sessions {
		fmt.Printf("Key: %v\t", k)
		fmt.Printf("Value: %v\n", v)
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get the session ID cookie
			sid, err := r.Cookie("sid")
			if err != nil {
				http.Error(w, "Forbidden. Please login to continue.", http.StatusForbidden)
				return
			}

			// Check if it is in the session storage.
			parsed, err := uuid.Parse(sid.Value)
			if err != nil {
				http.Error(w, "Forbidden. Please login to continue.", http.StatusForbidden)
				return
			}
			session, found := (*sessions)[parsed]

			// If not found, return with unauthorized error
			if !found {
				http.Error(w, "Forbidden. Please login to continue.", http.StatusForbidden)
				return
			}

			// If found, continue normal execution
			log.Printf("UserID: %v\n", session.UID)
			log.Printf("Role: %v\n", session.Role)
			next.ServeHTTP(w, r)
		})
	}
}
