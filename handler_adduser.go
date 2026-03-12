package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/jondatkins/Chirpy/internal/auth"
	"github.com/jondatkins/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerAddUser(w http.ResponseWriter, r *http.Request) {
	isDev := os.Getenv("PLATFORM")
	if isDev != "dev" {
		respondWithError(w, http.StatusForbidden, "Not allowed for non devs!!!", nil)
		return
	}
	type parameters struct {
		Password string `json:password`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	hashedPassword, err := auth.HashPassword(params.Password)
	if err != nil {
		fmt.Printf("Error hashing password for user %s", params.Email)
		return
	}
	user, err := cfg.db.CreateUser(r.Context(), database.CreateUserParams{
		params.Email,
		hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating user", err)
	}
	newUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
	fmt.Println("User created successfully:", user)
	respondWithJSON(w, http.StatusCreated, newUser)
}
