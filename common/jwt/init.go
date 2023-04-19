package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	AccessTokenKey []byte
	AdminTokenKey  []byte
}

func NewJWTManager(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) GenerateUserToken(id string, duration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		},
	})

	stringAccessToken := string(j.AccessTokenKey)

	userAccessToken := []byte(stringAccessToken)

	tokenString, err := token.SignedString(userAccessToken)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (j JWTManager) VerifyUserToken(tokenString string) (claims *UserClaims, err error) {
	stringAccessToken := string(j.AccessTokenKey)

	userAccessToken := []byte(stringAccessToken)

	tkn, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return userAccessToken, nil
	})

	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		err = errors.New("token invalid")
		return nil, err
	}

	return claims, nil
}
