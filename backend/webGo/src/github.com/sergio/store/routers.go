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
		j, _ := json.Marshal("Error in token check")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
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

	claims, err := common.GetTokenParsed(params["token"])
	if err == false {
		j, _ := json.Marshal("Error in token check")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	} else {
		id := GetIdByUsername(fmt.Sprintf("%v", claims["sub"]))
		data := GetUserOrders(id)

		j, _ := json.Marshal(data)
		fmt.Println(string(j))

		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}
func CreateOrder(w http.ResponseWriter, r *http.Request) {
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
	db_save := SaveOrder(id, params["amount"], params["currency"], params["price"])
	if db_save {
		j, _ := json.Marshal("Order OK")
		w.WriteHeader(http.StatusOK)
		w.Write(j)
		return
	} else {
		j, _ := json.Marshal("Database error")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
}
