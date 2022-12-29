package model

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Name       string `json:"name"`
	Authorized bool   `json:"authorized"`
	Exp        time.Time   `json:"exp"`
	jwt.StandardClaims
}

