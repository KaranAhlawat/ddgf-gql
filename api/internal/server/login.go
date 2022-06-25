package server

import (
	"ddgf-new/internal/repo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type Password struct {
	Password string `json:"pass"`
}

func LoginHandler(sessions *map[uuid.UUID]repo.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var body Password
		err := decoder.Decode(&body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid POST body")
			return
		}

		if body.Password != "deita" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Invalid Password.")
			return
		}
		id, err := uuid.Parse("fe1398ca-2feb-4cca-897f-f58e7b7aab3d")
		if err != nil {
			return
		}

		sid := uuid.New()

		(*sessions)[id] = repo.Session{
			Role: "USER",
			SID:  sid,
			UID:  id,
		}

		for k, v := range *sessions {
			fmt.Printf("Key: %v\t", k)
			fmt.Printf("Value: %v\n", v)
		}

		http.SetCookie(w, &http.Cookie{
			Name:  "sid",
			Value: sid.String(),
		})

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Logged in.")
	})
}
