package model

import (
	"github.com/golang-jwt/jwt/v4"
)

type UserClaims struct {
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type User struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
