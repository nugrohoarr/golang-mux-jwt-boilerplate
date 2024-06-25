package config

import (
	"os"

	"github.com/golang-jwt/jwt/v4"
)

var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

type JWTClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
