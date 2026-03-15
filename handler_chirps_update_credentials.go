package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jondatkins/Chirpy/internal/auth"
	"github.com/jondatkins/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerUpdateCredentials(w http.ResponseWriter, r *http.Request) {
	type response struct {
		User
		Token        string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	authHeader, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "No access token", err)
		return
	}
	userId, err := auth.ValidateJWT(authHeader, cfg.secretKey)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid JWT", err)
		return
	}
	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		fmt.Printf("Error hashing password for user %s", params.Email)
		return
	}
	user, err := cfg.db.UpdateEmailAndPassword(r.Context(), database.UpdateEmailAndPasswordParams{
		ID:             userId,
		Email:          params.Email,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error updating credentials ", err)
	}
	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
			Email:     user.Email,
		},
	})
}
