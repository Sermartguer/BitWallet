package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type Creds struct {
	Status      string
	APIKey      string
	AccountType string
	Email       string
	AuthToken   string
	IsLoggedIn  bool
}

func Login(w http.ResponseWriter, r *http.Request) {
	dat, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	log.Printf(params["username"])
	credentials := GetCredentials(params["username"], params["password"])

	out, _ := json.MarshalIndent(&credentials, "", "  ")
	fmt.Fprintf(w, string(out))

}
func idLogged(w http.ResponseWriter, r *http.Request) {
	//var jwtResponse jwtKey
	dat, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	log.Println(reflect.TypeOf(params["jwt"]))
	hasValidToken(params["jwt"])
}
func hasValidToken(jwtToken string) bool {
	ret := false
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return myLookupKey(), nil
	})

	if err == nil && token.Valid {
		ret = true
	}
	log.Println(ret)

	return ret
}
func GetCredentials(username, password string) Creds {
	credentials := Creds{
		Status:      "UNAUTHORIZED",
		APIKey:      "",
		AccountType: "",
		Email:       "",
		AuthToken:   "",
		IsLoggedIn:  false,
	}
	if (username == "admin") && (password == "admin") {
		credentials.Status = "OK"
		credentials.APIKey = "12345"
		credentials.AccountType = "admin"
		credentials.Email = "admin@example.com"
		credentials.IsLoggedIn = true
		// Now create a JWT for user
		// Create the token
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		// Set some claims
		claims["sub"] = username
		claims["iss"] = "example.com"
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		token.Claims = claims
		var err error
		credentials.AuthToken, err = token.SignedString([]byte("biscuits and gravy"))
		if err != nil {
			log.Println(err)
		}
	}
	return credentials
}

func myLookupKey() []byte {
	return []byte("biscuits and gravy")
}
