package main

import "testing"

func TestDirtyString(t *testing.T) {
	dirtyString := "This is a kerfuffle opinion I need to share with the world"
	cleanString := "This is a **** opinion I need to share with the world"
	resultString, _ := validateChirp(dirtyString)
	if cleanString != resultString {
		t.Errorf(`Expect resultString '%s' to equal cleanString '%s'`, resultString, cleanString)
	}
}

func TestCleanString(t *testing.T) {
	cleanString := "This is a kerfuffle! opinion I need to share with the world"
	resultString, _ := validateChirp(cleanString)

	if cleanString != resultString {
		t.Errorf(`Expect resultString '%s' to equal cleanString '%s'`, resultString, cleanString)
	}
}
