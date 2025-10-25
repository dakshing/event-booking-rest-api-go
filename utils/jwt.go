package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secret = "NOT_SO_SECRET_CHANGE_ME"

func GenerateTokenString(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(secret))
}

func ValidateTokenStringAndGetUserID(tokenString string) (int64, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token passed")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token passed")
	}

	return int64(claims["userId"].(float64)), nil
}
