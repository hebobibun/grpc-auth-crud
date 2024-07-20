package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

type JWTwrapper struct {
	SecretKey  string
	Issuer     string
	Expiration int64
}

type JWTclaims struct {
	jwt.StandardClaims
	Id    int64
	Email string
}

func (j *JWTwrapper) GenerateToken(id int64, email string) (string, error) {
	claims := JWTclaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    j.Issuer,
			ExpiresAt: time.Now().Add(time.Second * time.Duration(j.Expiration)).Unix(),
		},
		Id:    id,
		Email: email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(j.SecretKey))
}

func (j *JWTwrapper) ValidateToken(signedToken string) (*JWTclaims, error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTclaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*JWTclaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	if claims.ExpiresAt < time.Now().Unix() {
		return nil, errors.New("token expired")
	}

	return claims, nil
}
