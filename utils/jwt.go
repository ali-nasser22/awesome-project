package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretKey = "verysecretkey,hahahahahaha" // exposing secret key. What ?!!!!! 😉😉😉

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(tokenString string) error {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return secretKey, nil
	})
	if err != nil {
		return errors.New("invalid token")
	}
	token := parsedToken.Valid
	if !token {
		return errors.New("invalid token")
	}
	// just to know how those can be extracted
	/*
		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			return errors.New("invalid token")
		}
		email := claims["email"].(string)
		userId := claims["userId"].(int64)
	*/
	return nil
}
