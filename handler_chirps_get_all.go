package main

import (
	"net/http"
)

func (cfg *apiConfig) handlerChirpsGetAll(w http.ResponseWriter, r *http.Request) {
	// type parameters struct {
	// 	UserID uuid.UUID `json:"user_id"`
	// }
	// decoder := json.NewDecoder(r.Body)
	// params := parameters{}
	// err := decoder.Decode(&params)
	// if err != nil {
	// 	respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
	// 	return
	// }
	chirps, err := cfg.db.GetAllChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "error getting chirps", err)
		return
	}
	chirpResponses := make([]Chirp, len(chirps))
	for i, chirp := range chirps {
		chirpResponses[i] = Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		}
	}
	respondWithJSON(w, http.StatusOK, chirpResponses)
}
