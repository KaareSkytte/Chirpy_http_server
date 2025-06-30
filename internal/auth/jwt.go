package auth

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func MakeJWT(userID uuid.UUID, tokenSecret string, expiresIn time.Duration) (string, error) {
	currentTime := time.Now().UTC()
	expireTime := currentTime.Add(expiresIn)

	jwtCurrentTime := jwt.NewNumericDate(currentTime)
	jwtExpireTime := jwt.NewNumericDate(expireTime)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwtCurrentTime,
		ExpiresAt: jwtExpireTime,
		Subject:   userID.String(),
	})

	tokenString, err := token.SignedString([]byte(tokenSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString, tokenSecret string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(tokenSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return uuid.Nil, errors.New("invalid token claims")
	}

	userIDString := claims.Subject

	userID, err := uuid.Parse(userIDString)
	if err != nil {
		return uuid.Nil, err
	}

	return userID, nil
}

func GetBearerToken(headers http.Header) (string, error) {
	bearerToken := headers.Get("Authorization")
	if len(bearerToken) == 0 {
		return "", errors.New("authorization headers doesn't exist")
	}
	list := strings.Split(bearerToken, " ")

	if len(list) != 2 || list[0] != "Bearer" {
		return "", errors.New("invalid authorization format")
	}

	token := list[1]

	return token, nil
}
