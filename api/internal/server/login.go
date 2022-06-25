package server

import (
	"context"
	"ddgf-new/internal/middleware"
	"ddgf-new/internal/repo"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

type Password struct {
	Password string `json:"pass"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
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
		log.Println("LoginHandler: password verification failed")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Invalid Password.")
		return
	}
	// Get the redis conn
	rc := r.Context().Value(middleware.RedisConn).(*redis.Client)

	// Replace with actual GORM query
	uid, err := uuid.Parse("fe1398ca-2feb-4cca-897f-f58e7b7aab3d")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error retreiving user.")
		return
	}

	sid := uuid.New()

	session := repo.Session{
		Role: "USER",
		SID:  sid,
		UID:  uid,
	}

	ms, err := json.Marshal(&session)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error logging in. Please try again.", http.StatusInternalServerError)
		return
	}

	err = rc.Set(context.Background(), sid.String(), ms, 0).Err()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error logging in. Please try again.", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:  "ddgf_sid",
		Value: sid.String(),
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged in.")
}
