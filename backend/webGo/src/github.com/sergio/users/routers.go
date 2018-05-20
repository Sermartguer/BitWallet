package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../common"
	"github.com/asaskevich/govalidator"
	_ "github.com/go-sql-driver/mysql"
)

func Islogged(w http.ResponseWriter, r *http.Request) {
	//var jwtResponse jwtKey
	dat, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	common.HasValidToken(params["jwt"])
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body) // Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	userExist := SearchUser(params["username"])
	if !userExist {
		common.StatusBadError(w, r, "Username not found")
		return
	}
	passwordHash := GetPassword(params["username"])
	checkPass := CheckPasswordHash(params["password"], passwordHash)
	if !checkPass {
		LoginActivity(GetIdByUsername(params["username"]), params["ip"], "0")
		common.StatusBadError(w, r, "Password not match")
		return
	}
	fmt.Println("aqui")

	LoginActivity(GetIdByUsername(params["username"]), params["ip"], "1")
	credentials := common.GetCredentials(params["username"], params["password"], GetEmail(params["username"]))
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
		common.StatusBadError(w, r, "Username or email already exist")
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

func VerifyAccount(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	// Read the body of the POST request
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(dat, &params)
	exists := CheckIfIdExists(params["param"])
	fmt.Println(exists)
	if !exists {
		common.StatusBadError(w, r, "Error verify")
		return
	}
	ver := Verify(params["param"], params["pin"])
	if ver {
		w.WriteHeader(http.StatusOK)
	} else {
		common.StatusBadError(w, r, "Error on verify")
		return
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
	claims, errorToken := common.GetTokenParsed(params["token"])
	if errorToken {
		common.StatusBadError(w, r, "Error in token check")
		return
	}

	userID := GetIdByUsername(fmt.Sprintf("%v", claims["sub"]))
	updateCheck := Update(params["fistname"], params["surname"], userID)
	if updateCheck {
		w.WriteHeader(http.StatusOK)
	} else {
		common.StatusBadError(w, r, "Error on update")
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
		common.StatusBadError(w, r, "Error in token check")
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
			common.StatusBadError(w, r, "Not updated")
			return
		}
	} else {
		common.StatusBadError(w, r, "Passwords not match")
		return
	}
	w.WriteHeader(http.StatusOK)
}
