package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type LoginAtempt struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CustomClaim struct {
	Email    string `json:"email"`
	FullName string `json:"fullname"`
	NickName string `json:"nickname"`
	SkpdId   int    `json:"skpd_id"`
	jwt.RegisteredClaims
}
