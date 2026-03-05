package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// func handlerValidate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(http.StatusText(http.StatusOK)))
// }

func handlerValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Message string `json:"body"`
	}

	type ErrorMsg struct {
		Error string `json:"errorMsg"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
		w.WriteHeader(500)
		return
	}
	// params is a struct with data populated successfully
	// ...
	msgLength := len(params.Message)
	if msgLength > 140 {
		fmt.Printf("Message is too long:%v\n", msgLength)
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		errorJson := ErrorMsg{
			Error: "Chirp is too long",
		}
		errorData, err := json.Marshal(errorJson)
		if err != nil {
			fmt.Printf("Error marshalling error data:%v\n", err)
		}
		w.Write([]byte(errorData))
		return
	}
	stripDirtyWords(params.Message)
	type returnVals struct {
		Valid bool `json:"valid"`
	}
	respBody := returnVals{
		Valid: true,
	}
	dat, err := json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}

func stripDirtyWords() {
	type cleanMessage struct {
		CleanMsg string `json:"cleanMessage"`
	}
	if strings.Contains(strings.ToLower(params.Message), "foo") {
		respMessage := cleanMessage{
			CleanMsg: "Too much foof\n",
		}
		msg, err := json.Marshal(respMessage)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(msg)
	}
}
