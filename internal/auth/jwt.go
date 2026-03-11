package auth

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.RegisteredClaims{
			Issuer:    "chirpy-access",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiresIn)),
			Subject:   userID.String(),
		})
	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	claims := &jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method is HMAC (not RSA, ECDSA, etc.)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// Return the secret key as []byte
		return []byte(tokenSecret), nil
	})
	// 3. Handle parse errors (expired, malformed, invalid signature)
	if err != nil {
		return uuid.Nil, err
	}

	// 4. Check if token passed validation
	if !token.Valid {
		return uuid.Nil, fmt.Errorf("invalid token")
	}

	// 5. Extract user ID from claims.Subject and return
	userId, err := uuid.Parse(claims.Subject)
	if err != nil {
		return uuid.Nil, err
	}
	return userId, nil
}

// This function should look for the Authorization header
// in the headers parameter and return the TOKEN_STRING if
// it exists (stripping off the Bearer prefix and whitespace).
// If the header doesn't exist, return an error.
func GetBearerToken(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("missing auth header")
	}
	token := strings.TrimPrefix(authHeader, "Bearer")
	trimmed_token := strings.TrimSpace(token)
	fmt.Println("auth token is:%s", trimmed_token)
	if trimmed_token == authHeader {
		return "", errors.New("invalid auth header format")
	}
	return trimmed_token, nil
}
