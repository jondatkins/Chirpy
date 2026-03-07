package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func (cfg *apiConfig) handlerAddUser(w http.ResponseWriter, r *http.Request) {
	isDev := os.Getenv("PLATFORM")
	if isDev != "dev" {
		respondWithError(w, http.StatusForbidden, "Not allowed for non devs!!!", nil)
		return
	}
	type parameters struct {
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	// user, err := cfg.db.CreateUser(r.Context(), params.Email)
	user, err := cfg.db.CreateUser(r.Context(), params.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating user", err)
	}
	newUser := User{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Email:     user.Email,
	}
	// jsonUser, err := json.Marshal(new_user)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "Marshalling error", err)
	// }
	fmt.Println("User created successfully:", user)
	respondWithJSON(w, http.StatusCreated, newUser)
}
