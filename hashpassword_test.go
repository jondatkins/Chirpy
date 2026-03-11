package main

import (
	"testing"

	"github.com/jondatkins/Chirpy/internal/auth"
)

func TestHashPassword(t *testing.T) {
	password := "password1"
	hashedPassword, _ := auth.HashPassword(password)
	isMatch, _ := auth.CheckPasswordHash(password, hashedPassword)
	if !isMatch {
		t.Errorf(`Expect password '%s' to equal hashedPassword '%s'`, password, hashedPassword)
	}
}

func TestWrongHashPassword(t *testing.T) {
	password := "password1"
	hashedPassword := "Foo Bar"
	isMatch, _ := auth.CheckPasswordHash(password, hashedPassword)
	if isMatch {
		t.Errorf(`Expect password '%s' to NOT equal hashedPassword '%s'`, password, hashedPassword)
	}
}
