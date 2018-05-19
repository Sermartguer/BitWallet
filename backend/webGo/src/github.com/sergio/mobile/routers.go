package mobile

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func LoginByMobileID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)
	data := GetUserData(params["mobileID"])
	j, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func MobileUserBalances(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)
	data := GetUserBalance(params["mobileID"])
	j, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
