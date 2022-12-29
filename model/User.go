package model

import (	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	Name       string    `json:"name"`
	jwt.StandardClaims
}

