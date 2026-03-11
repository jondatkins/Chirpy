package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/jondatkins/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Password         string `json:password`
		Email            string `json:"email"`
		ExpiresInSeconds *int   `json:"expires_in_seconds"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	duration := time.Hour
	if params.ExpiresInSeconds != nil {
		client_duration := time.Duration(*params.ExpiresInSeconds) * time.Second
		if client_duration < time.Hour {
			duration = client_duration
		}
	}
	user, err := cfg.db.GetUserByEmail(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}

	isMatch, _ := auth.CheckPasswordHash(params.Password, user.HashedPassword.String)
	if !isMatch {
		respondWithError(w, http.StatusUnauthorized, "Incorrect email or password", err)
		return
	}
	jwtToken, err := auth.MakeJWT(user.ID, cfg.secretKey, duration)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error making JWT", err)
	}
	loggedInUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
		Token:     jwtToken,
	}
	fmt.Println("User logged in successfully:", user)
	respondWithJSON(w, http.StatusOK, loggedInUser)
}
