package config

import (
	"fmt"
	"os"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

var JWT_KEY []byte

// InitializeJWTKey initializes the JWT key from environment or a secure source
func InitializeJWTKey() error {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env file: %v", err)
	}

	// Get JWT_KEY from environment variables
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return fmt.Errorf("JWT_KEY not set in .env file")
	}

	// Assign to JWT_KEY
	JWT_KEY = []byte(jwtKey)
	return nil
}

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
