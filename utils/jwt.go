package utils

import (
	"time"
	"webhookserver/structs"

	"github.com/golang-jwt/jwt/v5"
)

func CreateJwt(origin string, seconds int) (string, error) {
	claims := jwt.MapClaims{
		"origin": origin,
		"iat":    time.Now().Unix(),
	}

	if seconds > 0 {
		claims["exp"] = time.Now().Add(time.Second * time.Duration(seconds)).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(structs.Env.JwtSecret)
}

func ValidateJwt(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", jwt.ErrInvalidKey
		}
		return structs.Env.JwtSecret, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", jwt.ErrInvalidKey
	}

	return claims["origin"].(string), nil
}
