package common

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
)

func GetTokenParsed(t string) (jwt.MapClaims, bool) {
	token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("do or do not there is no try"), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	} else {
		return nil, false
	}
}
