package overview

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"../common"
)

func GetUserGenericData(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	claims, err := common.GetTokenParsed(params["token"])
	if err == false {
		j, _ := json.Marshal("Error in token check")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	id := GetIdByUsername(fmt.Sprintf("%v", claims["sub"]))
	data := GetGenericData(id)
	UpdateBalance(fmt.Sprintf("%v", claims["sub"]), "DOGE")
	UpdateBalance(fmt.Sprintf("%v", claims["sub"]), "LTC")
	UpdateBalance(fmt.Sprintf("%v", claims["sub"]), "BTC")
	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func GetCoinPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	body := GetCurrencyPrice(os.Getenv(params["currency"] + "TST"))
	if string(body) == "Error" {
		j, _ := json.Marshal("Block.io error")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	claims, err := common.GetTokenParsed(params["token"])
	if err == false {
		j, _ := json.Marshal("Error in token check")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}

	userID := GetIdByUsername(fmt.Sprintf("%v", claims["sub"]))
	data := UserTransactions(userID)
	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func UpdateBalance(username string, currency string) {
	userID := GetIdByUsername(username)
	active := CheckAddress(userID, currency)
	if active {
		labelName := GetLabelByID(userID, currency)
		body := UpdateBalancesApi(currency, labelName)

		data := &ResponseGetBalance{
			Data: &Balance{},
		}
		json.Unmarshal(body, data)
		log.Println("Update")
		fmt.Println(string(body))
		log.Println(data.Data.Balance)
		UpdateBalanceTo(data.Data.Balance, userID, currency)
	}
}
