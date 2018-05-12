package common

import (
	"fmt"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//Creds used to store token
type Creds struct {
	AuthToken string
}

//GetCredentials generate token by user params, returns a Creds struct
func GetCredentials(username string, password string, email string) Creds {
	credentials := Creds{
		AuthToken: "",
	}
	// Now create a JWT for user
	// Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	// Set some claims
	claims["sub"] = username
	claims["iss"] = email
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token.Claims = claims
	var err error
	credentials.AuthToken, err = token.SignedString([]byte("do or do not there is no try"))
	if err != nil {
		log.Println(err)
	}
	return credentials
}

//MyLookupKey used to make token with key sentence
func MyLookupKey() []byte {
	return []byte("do or do not there is no try")
}

//HasValidToken check if token is valid, returns a check error
func HasValidToken(jwtToken string) bool {
	ret := false
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return MyLookupKey(), nil
	})
	if err == nil && token.Valid {
		ret = true
	}
	log.Println(ret)

	return ret
}

//GetUsernameByToken get username by an a token
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

//GetEmailByToken get email by an a token
func GetEmailByToken(token string) (string, bool) {
	var response string
	var errorToken bool
	claims, err := GetTokenParsed(token)
	if err == false {
		response = "Error in token check"
		errorToken = true
	} else {
		response = fmt.Sprintf("%v", claims["iss"])
		errorToken = false
	}
	return response, errorToken
}

//GetExpirationByToken get time expiration by an a token
func GetExpirationByToken(token string) (string, bool) {
	var response string
	var errorToken bool
	claims, err := GetTokenParsed(token)
	if err == false {
		response = "Error in token check"
		errorToken = true
	} else {
		response = fmt.Sprintf("%v", claims["exp"])
		errorToken = false
	}
	return response, errorToken
}
