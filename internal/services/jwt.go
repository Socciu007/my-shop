package services

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

type PayloadType struct {
	ID   string `json:"id"`
	Role string `json:"role"`
}

// generates a JWT access token for the given payload
func GenerateAccessToken(payload PayloadType, secretKey string) (string, error) {
	// Define the token expiration time (30s)
	expTime := time.Now().Add(30 * time.Second)

	// Create the JWT claims, which includes the payload and expiration time
	claims := &jwt.MapClaims{
        "id":   payload.ID,
        "role": payload.Role,
		"exp": expTime.Unix(),
    }

    // Create the JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and return the token string
	accessToken, err := token.SignedString([]byte(secretKey))
	if err!= nil {
        return "", err
    }

    return accessToken, nil
}

// generateRefreshToken generates a refresh token
func GenerateRefreshToken(payload PayloadType, secretKey string) (string, error) {
	// Define the token expiration time (365 days)
    expTime := time.Now().Add(365* 24 * time.Hour)

    // Create the JWT claims, which includes the payload and expiration time
    claims := &jwt.MapClaims{
        "id":   payload.ID,
        "role": payload.Role,
        "exp": expTime.Unix(),
    }

    // Create the JWT token
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

    // Sign and return the token string
    refreshToken, err := token.SignedString([]byte(secretKey))
    if err!= nil {
        return "", err
    }

    return refreshToken, nil
}