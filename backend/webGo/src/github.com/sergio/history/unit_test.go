package history

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

	r.HandleFunc("/api/loginHistory", GetLoginHistory).Methods("POST")

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return r
}

var (
	server    *httptest.Server
	reader    io.Reader //Ignore this for now
	usersUrl  string
	token     string
	tokenFail string
)

func init() {
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjYzNzQ5NjEsImlzcyI6InNlcm1hcnRndWVyQGdtYWlsLmNvbSIsInN1YiI6InNlcm1hcmd1ZXIifQ.DmSpLTdnuarRoyEjO9LACboZE_t3BnMOh0V-UjiSOVw"
	tokenFail = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MjYzNzQ5NjEsImlzcyI6InNlcm1hcnRndWVyQGdtYWlsLmNvbSIsInN1YiI6InNlcm1hcmd1ZXIifQ.DmSpLTdnuarRoyEjO9LACboZE_t3BnMOh0V-UjiSOV"
	server = httptest.NewServer(Handlers())                   //Creating new server with the user handlers
	usersUrl = fmt.Sprintf("%s/api/loginHistory", server.URL) //Grab the address for the API endpoint
}

/*
 API TEST - GetCoinPrice METHOD
*/
func TestLoginHistory(t *testing.T) {
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
func TestLoginHistoryFail(t *testing.T) {
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
