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
	db_save := SaveOrder(id, params["amount"], params["currency"], params["price"], params["currency_to"])
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
func BuyOrder(w http.ResponseWriter, r *http.Request) {
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
	//1. UpdateOrder
	update := SetPayment(params["restAmount"], params["payTo"], params["currencyO"])
	if !update {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//2.Comprar
	labelOrigen := GetLabelFromId(id, params["currencyO"])
	id_user := GetIdFromIdOrder(params["payTo"])
	labelDestino := GetLabelFromId(id_user, params["currencyO"])
	fmt.Println(labelOrigen)
	fmt.Println(labelDestino)
	if labelOrigen == "" {
		common.StatusBadError(w, r, "Error you don't have "+params["currencyO"]+" address")
		return
	} else if labelDestino == "" {
		common.StatusBadError(w, r, "Error you user order don't have "+params["currencyO"]+" address")
		return
	}

	//3.Vender
	labelOrigenV := GetLabelFromId(id, params["currencyD"])
	id_userV := GetIdFromIdOrder(params["payTo"])
	labelDestinoV := GetLabelFromId(id_userV, params["currencyD"])
	fmt.Println(labelOrigenV)
	fmt.Println(labelDestinoV)
	if labelOrigenV == "" {
		common.StatusBadError(w, r, "Error you don't have "+params["currencyO"]+" address")
		return
	} else if labelDestinoV == "" {
		common.StatusBadError(w, r, "Error you user order don't have "+params["currencyO"]+" address")
		return
	}

	recive := SendBalanceByAddress(params["currencyO"], params["totalToPay"], labelDestino, labelOrigen)
	send := SendBalanceByAddress(params["currencyD"], params["totalToRecive"], labelOrigenV, labelDestinoV)
	fmt.Println(string(recive))
	fmt.Println(string(send))
	//labelFrom := GetLabelFromId(id, params["currencyO"])
	/*j, _ := json.Marshal(labelFrom)
	fmt.Println(string(j))*/

	w.WriteHeader(http.StatusOK)
	//w.Write(j)
	//SendBalanceByAddress()
}
