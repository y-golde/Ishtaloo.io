package cookieUtils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

func decodeToken(tokenString string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return claims, err
	}
	return claims, nil

}
