package common

import (
	"fmt"
)

type TokenUsername struct {
	Username string `json:"sub"`
	Error    string `json:"error"`
}

func GetUsernameByToken(token string) (string, bool) {
	var response string
	var errorToken bool
	claims, err := GetTokenParsed(token)
	if err == false {
		response = "Error in token check"
		errorToken = true
	} else {
		response = fmt.Sprintf("%v", claims["sub"])
		errorToken = false
	}
	return response, errorToken
}
