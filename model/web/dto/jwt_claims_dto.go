package dto

import "github.com/golang-jwt/jwt"

type JwtClaimsDto struct {
	jwt.StandardClaims
	UserId int64  `json:"userId"`
	Email  string `json:"email"`
	Name   string `json:"name"`
}
