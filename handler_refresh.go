package main

import (
	"net/http"
	"time"

	"github.com/jondatkins/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "No auth header", err)
		return
	}
	user, err := cfg.db.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error getting user with token", err)
	}
	jwtToken, err := auth.MakeJWT(
		user.ID,
		cfg.secretKey,
		time.Hour,
	)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't validate token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, response{
		Token: jwtToken,
	})
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "Couldn't find token", err)
		return
	}
	_, err = cfg.db.RevokeRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not revoke refresh token", err)
		return
	}
	// respondWithJSON(w, http.StatusNoContent, nil)
	w.WriteHeader(http.StatusNoContent)
}
