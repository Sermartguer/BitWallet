package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"../common"
	"../dashboard"
	"github.com/asaskevich/govalidator"
	jwt "github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
)

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
		LoginActivity(GetIdByUsername(params["username"]), params["ip"], "0")
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Password not match")
		w.Write(j)
		return
	}
	dashboard.UpdateBalance(params["username"], "DOGE")
	dashboard.UpdateBalance(params["username"], "BTC")
	dashboard.UpdateBalance(params["username"], "LTC")
	LoginActivity(GetIdByUsername(params["username"]), params["ip"], "1")
	credentials := GetCredentials(params["username"], params["password"], GetEmail(params["username"]))
	out, _ := json.MarshalIndent(&credentials, "", "  ")
	fmt.Fprintf(w, string(out))
}
func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	govalidator.SetFieldsRequiredByDefault(true)
	dat, _ := ioutil.ReadAll(r.Body)

	//Validacio de camps
	userData := ValidateParams(dat)
	if !userData.Error {
		log.Printf(userData.TextError)
		j, _ := json.Marshal(userData.TextError)
		w.WriteHeader(http.StatusUnprocessableEntity)
		w.Write(j)
		return
	}
	//Comparacio entre els camps de contrasenya
	match := CheckPasswordHash(userData.Password2, userData.Password)
	if !match {
		w.WriteHeader(http.StatusUnprocessableEntity)
		j, _ := json.Marshal("Passwords does not match")
		w.Write(j)
		return
	}
	//Comprovacio si el usuari ja existeix en la BD
	userExist := CheckUser(userData)
	if userExist {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Username or email already exist")
		w.Write(j)
		return
	}

	//Registrar l'usuari en la BD
	saved := SaveUser(userData, "mob-"+userData.ID)
	if !saved {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	//Enviar email
	common.SendMail("check", userData.ID, userData.Username, userData.Email)
	//Enviar al frontend
	log.Printf("Saved on DB")
	w.WriteHeader(http.StatusCreated)
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
func VerifyAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	// Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	ver := Verify(params["param"], params["pin"])
	if ver {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Error on verify")
		w.Write(j)
	}
	log.Println(params["param"])
}
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	// Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	log.Println(params["fistname"])
	claims, err := common.GetTokenParsed(params["token"])
	if err == false {
		j, _ := json.Marshal("Error in token check")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}

	userID := GetIdByUsername(fmt.Sprintf("%v", claims["sub"]))
	updateCheck := Update(params["fistname"], params["surname"], userID)
	if updateCheck {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		j, _ := json.Marshal("Error on update")
		w.Write(j)
	}
}
func GetAccountProfile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	// Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	username, errorToken := common.GetUsernameByToken(params["token"])
	if errorToken {
		j, _ := json.Marshal("Error in token check")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	userID := GetIdByUsername(username)
	data := GetAccount(userID)
	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func RecoverPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	// Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	id := GetIdByEmail(params["email"])
	common.SendMail("newPassword", id, params["email"], params["email"])
}
func NewPassword(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	// Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)

	passwordHash, _ := HashPassword(params["password"])
	passCheck := CheckPasswordHash(params["repassword"], passwordHash)
	if passCheck {
		log.Println("OK")
		passDB := NewAccountPassword(passwordHash, params["id"])
		if passDB {
			j, _ := json.Marshal("OK")
			w.WriteHeader(http.StatusOK)
			w.Write(j)
			return
		} else {
			j, _ := json.Marshal("Not updated")
			w.WriteHeader(http.StatusBadRequest)
			w.Write(j)
			return
		}
	} else {
		j, _ := json.Marshal("Passwords not match")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	w.WriteHeader(http.StatusOK)
}
