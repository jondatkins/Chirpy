package auth

import (
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
)

// "github.com/jondatkins/Chirpy/internal/auth"

func TestCreateJWT(t *testing.T) {
	// id := uuid.New()
	id, _ := uuid.Parse("123")
	secret := "nomoresecrets"
	dur := time.Hour
	_, err := MakeJWT(id, secret, dur)
	if err != nil {
		t.Errorf(`Error making jwt with secret '%v'`, err)
	}
}

func TestValidateJWT(t *testing.T) {
	// id := uuid.New()
	id, _ := uuid.Parse("123")
	secret := "nomoresecrets"
	dur := time.Hour
	tokenString, err := MakeJWT(id, secret, dur)
	if err != nil {
		t.Errorf(`Error making jwt with secret '%v'`, err)
	}
	_, err = ValidateJWT(tokenString, secret)
	if err != nil {
		t.Errorf(`Error validating jwt with secret '%v'`, err)
	}
}

func TestGetBearerToken(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer mytoken123")
	_, err := GetBearerToken(headers)
	if err != nil {
		t.Errorf(`Error getting bearer token'%v'`, err)
	}
}

func TestGetBearerTokenNoAuth(t *testing.T) {
	headers := http.Header{}
	headers.Set("NotAuthorization", "Bearer mytoken123")
	_, err := GetBearerToken(headers)
	if err == nil {
		t.Errorf(`Should not be able to get bearer token'%v'`, err)
	}
}
