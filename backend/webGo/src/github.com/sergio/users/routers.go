package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

type Creds struct {
	AuthToken string
}
type User struct {
	ID          string `json:"id" valid:"-"`
	Username    string `json:"username" valid:"required~Username is blank"`
	Email       string `json:"email" valid:"email"`
	Password    string `json:"password" valid:"length(5|10)"`
	CreatedAt   string `json:"create_at" valid:"-"`
	AccountType string `json:"acc_type" valid:"optional"`
}
type Error struct {
	TextError string `json:"error"`
}

func Islogged(w http.ResponseWriter, r *http.Request) {
	//var jwtResponse jwtKey
	dat, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	log.Println(reflect.TypeOf(params["jwt"]))
	hasValidToken(params["jwt"])
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	userExist := SearchUser(params["username"])
	if !userExist {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Username not found")
		w.Write(j)
		return
	}
	passwordHash := GetPassword(params["username"])
	checkPass := CheckPasswordHash(params["password"], passwordHash)
	if !checkPass {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Password not match")
		w.Write(j)
		return
	}

	credentials := GetCredentials(params["username"], params["password"], GetEmail(params["username"]))
	out, _ := json.MarshalIndent(&credentials, "", "  ")
	fmt.Fprintf(w, string(out))
}
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	govalidator.SetFieldsRequiredByDefault(true)
	dat, _ := ioutil.ReadAll(r.Body)
	userData := ValidateParams(dat)
	var ok bool
	ok = true
	if !userData.Error {
		log.Printf(userData.TextError)
		j, _ := json.Marshal(userData.TextError)
		ok = false
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(j)
		return
	}

	match := CheckPasswordHash(userData.Password2, userData.Password)
	if !match {
		ok = false
		w.WriteHeader(http.StatusUnprocessableEntity)
		j, _ := json.Marshal("Passwords does not match")
		w.Write(j)
		return
	}
	userExist := CheckUser(userData)
	if userExist {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Username or email already exist")
		w.Write(j)
		return
	}
	saved := SaveUser(userData)
	if !saved {
		ok = false
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if ok {
		log.Printf("Saved on DB")
		w.WriteHeader(http.StatusCreated)
		return
	}
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

func myLookupKey() []byte {
	return []byte("do or do not there is no try")
}
