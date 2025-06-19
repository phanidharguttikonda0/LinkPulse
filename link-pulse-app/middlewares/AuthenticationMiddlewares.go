package middlewares

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func AuthorizationCheck(secret string, authorizationHeader string) (jwt.MapClaims, error) {
	jwtSecret := []byte(secret)
	token, err := jwt.Parse(authorizationHeader, func(token *jwt.Token) (interface{}, error) {
		// validating signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return jwtSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims")
	}

	return claims, nil
}

func CreateAuthorizationHeader(secret string, userId int, username string) (string, error) {
	jwtSecret := []byte(secret)
	claims := jwt.MapClaims{
		"username": username,
		"userId":   userId,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // expires in 24 hours
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}
