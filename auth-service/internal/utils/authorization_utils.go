package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type CustomClaims struct {
	jwt.RegisteredClaims

	Id string `json:"id"`
}

var mySecret = []byte(os.Getenv("JWT_SECRET"))

func NewToken(id string) string {

	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * 30 * time.Minute)),
		},
		Id: id,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, _ := jwtToken.SignedString(mySecret)

	return token
}

func VerifyToken(tokenString string) error {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error verifying algorithm")
		}

		return mySecret, nil
	})

	if err != nil || !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func ExtractTokenId(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("error verifying algorithm")
		}

		return mySecret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return jwt.MapClaims{}, err
}
