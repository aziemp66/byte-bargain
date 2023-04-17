package jwt

import "github.com/golang-jwt/jwt/v4"

type UserClaims struct {
	ID string `json:"id"`
	jwt.RegisteredClaims
}
