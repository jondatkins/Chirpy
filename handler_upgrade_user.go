package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/jondatkins/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerUpgradeUser(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	polkaKey := os.Getenv("POLKA_KEY")
	if polkaKey == "" {
		log.Fatal("POLKA_KEY must be set")
	}
	authHeader, err := auth.GetAPIKey(r.Header)
	if err != nil || polkaKey != authHeader {
		respondWithError(w, http.StatusUnauthorized, "No auth header", err)
		return
	}
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil || params.Event != "user.upgraded" {
		respondWithError(w, http.StatusNoContent, "couldn't decode parameters", err)
		return
	}

	_, err = cfg.db.UpgradeUser(r.Context(), params.Data.UserID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get user for id", err)
		return
	}

	respondWithJSON(w, http.StatusNoContent, nil)
}
