package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../common"
)

func GetOrdersEndpoint(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	_, err := common.GetTokenParsed(params["token"])
	if err == false {
		common.StatusBadError(w, r, "Error in token check")
		return
	} else {
		data := GetOrders()

		j, _ := json.Marshal(data)
		fmt.Println(string(j))

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

func GetOrdersUserEndpoint(w http.ResponseWriter, r *http.Request) {
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
	data := GetUserOrders(id)

	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
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
	db_save := SaveOrder(id, params["amount"], params["currency"], params["price"])
	if db_save {
		j, _ := json.Marshal("Order OK")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	} else {
		common.StatusBadError(w, r, "Database error")
		return
	}
}
func GetBalanceOnOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	data := GetBalance(id, params["currency"])

	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)

}
