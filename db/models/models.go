package models

import (
	"csrf/randomstrings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type User struct {
	UUID string `json:"uuid"`
	Username string `json:"user" bson:"user"`
	PasswordHash string `json:"password"`
	Role string
}

// https://tools.ietf.org/html/rfc7519
type TokenClaims struct {
	jwt.StandardClaims
	Role string `json:"role"`
	Csrf string `json:"csrf"`
}

type RefreshToken struct {
    JTI string
    Status string // You can use this field to manage token status, e.g., "valid", "used", "expired," etc.
}


const RefreshTokenValidTime = time.Hour * 72
const AuthTokenValidTime = time.Minute * 15

func GenerateCSRFSecret() (string, error) {
	return randomstrings.GenerateRandomString(32)
}
