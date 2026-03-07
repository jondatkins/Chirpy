package main

import (
	"fmt"
	"strings"
	"unicode"
)

// func handlerChirps(w http.ResponseWriter, r *http.Request) {
// 	type parameters struct {
// 		Body   string `json:"body"`
// 		UserID string `json:"user_id"`
// 	}
// 	// type returnVals struct {
// 	// 	Valid bool `json:"valid"`
// 	// }
// 	type returnVals struct {
// 		CleanedBody string `json:"cleaned_body"`
// 	}
//
// 	decoder := json.NewDecoder(r.Body)
// 	params := parameters{}
// 	err := decoder.Decode(&params)
// 	if err != nil {
// 		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
// 		return
// 	}
//
// 	const maxChirpLength = 140
// 	if len(params.Body) > maxChirpLength {
// 		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
// 		return
// 	}
// 	msg, _ := stripDirtyWords(params.Body)
//
// 	respondWithJSON(w, http.StatusOK, returnVals{
// 		CleanedBody: msg,
// 	})
// }

func stripDirtyWords(tweet string) (string, bool) {
	dirtyWords := []string{"kerfuffle", "sharbert", "fornax"}
	tweetWords := strings.Split(tweet, " ")
	type cleanMessage struct {
		CleanMsg string `json:"cleanMessage"`
	}
	returnString := ""
	isDirty := false
	for i, tweetWord := range tweetWords {
		for _, dirtyWord := range dirtyWords {
			lastRune := []rune(tweetWord)[len([]rune(tweetWord))-1]
			if strings.ToLower(tweetWord) == dirtyWord && !unicode.IsPunct(lastRune) {
				isDirty = true
				break
			} else {
				isDirty = false
			}

		}
		if isDirty {
			returnString = returnString + " ****"
		} else {
			if i == 0 {
				returnString = returnString + tweetWord
			} else {
				returnString = returnString + " " + tweetWord
			}
		}
	}
	fmt.Println(returnString)
	return returnString, isDirty
}
