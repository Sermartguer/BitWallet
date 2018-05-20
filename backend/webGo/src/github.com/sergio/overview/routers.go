package overview

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"../common"
)

func GetCoinPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	body := GetCurrencyPrice(os.Getenv(params["currency"] + "TST"))
	if string(body) == "Error" {
		common.StatusBadError(w, r, "Block.io error")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
func GetUserGenericData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	username, errorToken := common.GetUsernameByToken(params["token"])
	if errorToken {
		common.StatusBadError(w, r, "Error in token check")
		return
	}
	id := GetIdByUsername(username)
	data := GetGenericData(id)
	UpdateBalance(username, "DOGE")
	UpdateBalance(username, "LTC")
	UpdateBalance(username, "BTC")
	j, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	username, errorToken := common.GetUsernameByToken(params["token"])
	if errorToken {
		common.StatusBadError(w, r, "Error in token check")
		return
	}

	userID := GetIdByUsername(username)
	data := UserTransactions(userID)
	j, _ := json.Marshal(data)

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func UpdateBalance(username string, currency string) {
	fmt.Println("ye")
	userID := GetIdByUsername(username)
	fmt.Println("Address")
	active := CheckAddress(userID, currency)
	if active {
		labelName := GetLabelByID(userID, currency)
		body := UpdateBalancesApi(currency, labelName)
		fmt.Println(string(body))
		data := &ResponseGetBalance{
			Data: &Balance{},
		}
		json.Unmarshal(body, data)
		fmt.Println("Update")
		UpdateBalanceTo(data.Data.Balance, userID, currency)
	}
}
