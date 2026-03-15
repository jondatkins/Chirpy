package main

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/jondatkins/Chirpy/internal/auth"
	"github.com/jondatkins/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {
	chirpId, err := uuid.Parse(r.PathValue("chirpId"))
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Need 'chirpId' id", err)
		return
	}
	authHeader, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "No auth header", err)
		return
	}
	user_id, err := auth.ValidateJWT(authHeader, cfg.secretKey)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Invalid JWT", err)
		return
	}

	chirp, err := cfg.db.GetChirpById(r.Context(), chirpId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "error getting chirp for id"+chirpId.String(), err)
		return
	}
	if chirp.UserID != user_id {
		respondWithError(w, http.StatusForbidden, "User not authorized", err)
		return
	}
	err = cfg.db.DeleteChirp(r.Context(), database.DeleteChirpParams{
		ID:     chirpId,
		UserID: user_id,
	})
	if err != nil {
		respondWithError(w, http.StatusNotFound, "error deleting chirp", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
