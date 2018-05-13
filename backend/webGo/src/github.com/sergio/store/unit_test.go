package store

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
	r.HandleFunc("/api/getOrders", GetOrdersEndpoint).Methods("POST")
	r.HandleFunc("/api/getUserOrders", GetOrdersUserEndpoint).Methods("POST")
	r.HandleFunc("/api/saveOrder", CreateOrder).Methods("POST")

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
	server = httptest.NewServer(Handlers())                       //Creating new server with the user handlers
	usersUrl = fmt.Sprintf("%s/api/getOrders", server.URL)        //Grab the address for the API endpoint
	userDataUrl = fmt.Sprintf("%s/api/getUserOrders", server.URL) //Grab the address for the API endpoint
	userTransUrl = fmt.Sprintf("%s/api/saveOrder", server.URL)    //Grab the address for the API endpoint
}

/*
	API TEST - GetOrdersEndpoint METHOD
*/
func TestGetOrders(t *testing.T) {
	userJson := `{"token":"` + token + `"}`
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
func TestGetOrdersFail(t *testing.T) {
	userJson := `{"token":"` + tokenFail + `"}`
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
func TestGetUserOrders(t *testing.T) {
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
func TestGetUsersOrdersFail(t *testing.T) {
	userJson := `{"token":"` + tokenFail + `"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userDataUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 400 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestGetCreateOrders(t *testing.T) {
	userJson := `{
		"token":"` + token + `",
		"amount":"1055",
		"price":"0.01",
		"currency":"DOGE"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userTransUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
func TestGetCreateOrdersFail(t *testing.T) {
	userJson := `{
		"token":"` + tokenFail + `",
		"amount":"1055",
		"price":"0.01",
		"currency":"DOGE"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userTransUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 400 {
		t.Errorf("Success expected: %d", res.StatusCode)
	}
}
