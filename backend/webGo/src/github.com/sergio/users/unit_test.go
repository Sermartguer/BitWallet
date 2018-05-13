package users

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func Handlers() *mux.Router {
	r := mux.NewRouter()

	//SEND API ENDPOINTS
	r.HandleFunc("/api/login", Login).Methods("POST")
	r.HandleFunc("/api/isLoged", Islogged).Methods("POST")
	r.HandleFunc("/api/register", Register).Methods("POST")
	r.HandleFunc("/api/verifyAccount", VerifyAccount).Methods("POST")
	r.HandleFunc("/api/updateProfile", UpdateProfile).Methods("POST")
	r.HandleFunc("/api/getAccountProfile", GetAccountProfile).Methods("POST")
	r.HandleFunc("/api/recoverPassword", RecoverPassword).Methods("POST")
	r.HandleFunc("/api/newPassword", NewPassword).Methods("POST")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return r
}

var (
	server         *httptest.Server
	reader         io.Reader //Ignore this for now
	usersUrl       string
	userDataUrl    string
	userTransUrl   string
	userAccountUrl string
	token          string
	tokenFail      string
)

func init() {
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjYzNzQ5NjEsImlzcyI6InNlcm1hcnRndWVyQGdtYWlsLmNvbSIsInN1YiI6InNlcm1hcmd1ZXIifQ.DmSpLTdnuarRoyEjO9LACboZE_t3BnMOh0V-UjiSOVw"
	tokenFail = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjYzNzQ5NjEsImlzcyI6InNlcm1hcnRndWVyQGdtYWlsLmNvbSIsInN1YiI6InNlcm1hcmd1ZXIifQ.DmSpLTdnuarRoyEjO9LACboZE_t3BnMOh0V-UjiSOV"
	server = httptest.NewServer(Handlers())                   //Creating new server with the user handlers
	usersUrl = fmt.Sprintf("%s/api/login", server.URL)        //Grab the address for the API endpoint
	userDataUrl = fmt.Sprintf("%s/api/isLoged", server.URL)   //Grab the address for the API endpoint
	userTransUrl = fmt.Sprintf("%s/api/register", server.URL) //Grab the address for the API endpoint
	userAccountUrl = fmt.Sprintf("%s/api/getAccountProfile", server.URL)
}

/*
	API TEST - GetOrdersEndpoint METHOD
*/
func TestIsLoged(t *testing.T) {
	userJson := `{"token":"` + token + `"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userDataUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestIsLogedFail(t *testing.T) {
	userJson := `{"token":"` + tokenFail + `"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userDataUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestLogin(t *testing.T) {
	userJson := `{"username": "sermarguer", "password": "111111","ip":"123"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", usersUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestLoginFail(t *testing.T) {
	userJson := `{"username": "sermarguer3", "password": "111111","ip":"123"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", usersUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 400 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestGetAccount(t *testing.T) {
	userJson := `{"token":"` + token + `"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userAccountUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestGetAccountFail(t *testing.T) {
	userJson := `{"token":"` + tokenFail + `"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userAccountUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 400 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
