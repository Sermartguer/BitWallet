package overview

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

	r.HandleFunc("/api/getCurrencyPrice", GetCoinPrice).Methods("POST")
	r.HandleFunc("/api/getUserData", GetUserGenericData).Methods("POST")
	r.HandleFunc("/api/getUserTrans", GetTransactions).Methods("POST")

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
	usersUrl = fmt.Sprintf("%s/api/getCurrencyPrice", server.URL) //Grab the address for the API endpoint
	userDataUrl = fmt.Sprintf("%s/api/getUserData", server.URL)   //Grab the address for the API endpoint
	userTransUrl = fmt.Sprintf("%s/api/getUserTrans", server.URL) //Grab the address for the API endpoint
}

/*
 API TEST - GetCoinPrice METHOD
*/
func TestCoinPriceBTC(t *testing.T) {
	userJson := `{"currency":"BTC"}`
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
func TestCoinPriceLTC(t *testing.T) {
	userJson := `{"currency":"LTC"}`
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
func TestCoinPriceDOGE(t *testing.T) {
	userJson := `{"currency":"DOGE"}`
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
func TestCoinPriceStatusSuccess(t *testing.T) {
	userJson := `{"currency":"DOGE"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", usersUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	} else {
		dat, _ := ioutil.ReadAll(res.Body)
		var params map[string]string
		json.Unmarshal(dat, &params)
		if params["status"] != "success" {
			t.Errorf("Success expected: success")
		}
	}
}
func TestCoinPriceStatusFail(t *testing.T) {
	userJson := `{"currency":"DOGEa"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", usersUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 200 {
		t.Errorf("Success expected: %d", res.StatusCode)
	} else {
		dat, _ := ioutil.ReadAll(res.Body)
		var params map[string]string
		json.Unmarshal(dat, &params)
		if params["status"] != "fail" {
			t.Errorf("Success expected: fail")
		}
	}
}

/*
API TEST - GetUserGenericData METHOD
*/
func TestGetGenericData(t *testing.T) {
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
func TestGetGenericDataFail(t *testing.T) {
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
func TestGetTransactionsEndpoint(t *testing.T) {
	userJson := `{"token":"` + token + `"}`
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

func TestGetTransactionsEndpointFail(t *testing.T) {
	userJson := `{"token":"` + tokenFail + `"}`
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
