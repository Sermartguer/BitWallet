package send

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

	//SEND API ENDPOINTS
	r.HandleFunc("/api/getAddresses", GetUserAddresses).Methods("POST")
	r.HandleFunc("/api/getNewAddress", GetNewAddress).Methods("POST")
	r.HandleFunc("/api/sendLocal", SendLocal).Methods("POST")
	r.HandleFunc("/api/sendExternal", SendExternal).Methods("POST")
	r.HandleFunc("/api/getFee", GetNetworkFee).Methods("POST")

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
	usersUrl = fmt.Sprintf("%s/api/getAddresses", server.URL)     //Grab the address for the API endpoint
	userDataUrl = fmt.Sprintf("%s/api/getFee", server.URL)        //Grab the address for the API endpoint
	userTransUrl = fmt.Sprintf("%s/api/getUserTrans", server.URL) //Grab the address for the API endpoint
}

/*
	API TEST - GetUserAddresses METHOD
*/
func TestGetAddresses(t *testing.T) {
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
func TestGetAddressesFail(t *testing.T) {
	userJson := `{"token":"` + tokenFail + `"}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", usersUrl, reader)
	//res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetUserAddresses)
	handler.ServeHTTP(rr, req)
	if rr.Body.String() != `"Error in token check"` {
		t.Errorf("Expected Error in token check")
	}

}

/*
	API TEST - GetNetworkFee METHOD
*/
func TestGetNetWorkFee(t *testing.T) {
	userJson := `{
		"currency":"DOGE",
		"token":"` + token + `",
		"amount":"20",
		"address":"2NA4s6PfcF6LYekVWyiSenQinKisX5bFHZf"
	}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userDataUrl, reader)
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
func TestGetNetWorkFeeFail(t *testing.T) {
	userJson := `{
		"currency":"DOGEa",
		"token":"` + token + `",
		"amount":"20",
		"address":"2NA4s6PfcF6LYekVWyiSenQinKisX5bFHZf"
	}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userDataUrl, reader)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}
	if res.StatusCode != 400 {
		t.Errorf("Success expected: %d", res.StatusCode)
	} else {
		dat, _ := ioutil.ReadAll(res.Body)
		var params map[string]string
		json.Unmarshal(dat, &params)
		if params["status"] != "fail" {
			t.Errorf("Success expected: success")
		}
	}
}
func TestGetNetWorkFeeResponse(t *testing.T) {
	userJson := `{
		"currency":"DOGE",
		"token":"` + token + `",
		"amount":"20",
		"address":"2NA4s6PfcF6LYekVWyiSenQinKisX5bFHZf"
	}`
	reader = strings.NewReader(userJson)
	req, err := http.NewRequest("POST", userDataUrl, reader)
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
		if params["network_fee"] != "2.00000000" {
			t.Errorf("Success expected: 2.00000000")
		}
	}
}
