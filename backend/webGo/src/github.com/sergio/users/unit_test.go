package users

import (
	"fmt"
	"io"
	"log"
	"net/http/httptest"

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
	server       *httptest.Server
	reader       io.Reader //Ignore this for now
	usersUrl     string
	userDataUrl  string
	userTransUrl string
	token        string
	tokenFail    string
)

func init() {
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjYzNzQ5NjEsImlzcyI6InNlcm1hcnRndWVyQGdtYWlsLmNvbSIsInN1YiI6InNlcm1hcmd1ZXIifQ.DmSpLTdnuarRoyEjO9LACboZE_t3BnMOh0V-UjiSOVw"
	tokenFail = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjYzNzQ5NjEsImlzcyI6InNlcm1hcnRndWVyQGdtYWlsLmNvbSIsInN1YiI6InNlcm1hcmd1ZXIifQ.DmSpLTdnuarRoyEjO9LACboZE_t3BnMOh0V-UjiSOV"
	server = httptest.NewServer(Handlers())                   //Creating new server with the user handlers
	usersUrl = fmt.Sprintf("%s/api/login", server.URL)        //Grab the address for the API endpoint
	userDataUrl = fmt.Sprintf("%s/api/isLoged", server.URL)   //Grab the address for the API endpoint
	userTransUrl = fmt.Sprintf("%s/api/register", server.URL) //Grab the address for the API endpoint
}

/*
	API TEST - GetOrdersEndpoint METHOD
*/
