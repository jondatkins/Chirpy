package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"github.com/jondatkins/Chirpy/internal/database"
)

func (cfg *apiConfig) handlerAddChirps(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string `json:"body"`
		UserID string `json:"user_id"`
	}

	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}
	const maxChirpLength = 140
	msgLength := len(params.Body)

	if msgLength > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", err)
	}
	badwords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	_, isClean := getCleanedBody(params.Body, badwords)
	if !isClean {
		respondWithError(w, http.StatusNotAcceptable, "Bad words found", nil)
	}
	// chirp, err := cfg.db.CreateChirp(r.Context(), params.Body, params.UserID)
	chirp, err := cfg.db.CreateChirp(r.Context(), database.CreateChirpParams{
		Body:   params.Body,
		UserID: uuid.MustParse(params.UserID),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error creating chirp", err)
	}
	newChirp := Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	}
	fmt.Println("Chirp created successfully:", newChirp)
	respondWithJSON(w, http.StatusCreated, newChirp)
	// respondWithJSON(w, http.StatusOK, returnVals{
	// 	CleanedBody: cleaned,
	// })
}

func getCleanedBody(body string, badWords map[string]struct{}) (string, bool) {
	words := strings.Split(body, " ")
	isClean := true
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
			isClean = false
		}
	}
	cleaned := strings.Join(words, " ")
	return cleaned, isClean
}
