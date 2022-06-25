package server

import (
	"context"
	"crypto/rand"
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
	var body Password
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, "Invalid POST body", http.StatusBadRequest)
		return
	}

	if body.Password != "deita" {
		log.Println("LoginHandler: password verification failed")
		http.Error(w, "Invalid password.", http.StatusBadRequest)
		return
	}
	// Get the redis conn
	rc := r.Context().Value(middleware.RedisConn).(*redis.Client)

	// Replace with actual GORM query
	uid, err := uuid.Parse("fe1398ca-2feb-4cca-897f-f58e7b7aab3d")
	if err != nil {
		log.Println(err)
		http.Error(w, "Error retreiving user.", http.StatusInternalServerError)
		return
	}

	// Create a 64 byte sessin ID
	sid := make([]byte, 64)
	_, err = rand.Read(sid)
	fmt.Printf("%x\n", sid)
	if err != nil {
		log.Println(err)
		http.Error(w, "Unable to login.", http.StatusInternalServerError)
		return
	}

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

	err = rc.Set(context.Background(), fmt.Sprintf("%x", sid), ms, 0).Err()
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error logging in. Please try again.", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   "ddgf_sid",
		Value:  fmt.Sprintf("%x", sid),
		MaxAge: 60 * 60 * 24 * 7,
	})

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Logged in.")
}
