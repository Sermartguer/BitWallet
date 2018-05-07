package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"unsafe"

	"../common"
)

type NewAddress struct {
	Status string   `json:"status"`
	Data   *Address `json:"data"`
}
type Address struct {
	Address string `json:"address"`
}
type ResponseGetBalance struct {
	Status string   `json:"status"`
	Data   *Balance `json:"data"`
}
type Balance struct {
	Balance string `json:"available_balance"`
}
type ResponseSendLocal struct {
	Status string                 `json:"status"`
	Data   *ResponseSendLocalData `json:"data"`
}
type ResponseSendLocalData struct {
	Network      string `json:"network"`
	TXID         string `json:"txid"`
	AWithdraw    string `json:"amount_withdrawn"`
	ASent        string `json:"amount_sent"`
	NetWorkFee   string `json:"network_fee"`
	BlockIOFee   string `json:"blockio_fee"`
	ErrorMessage string `json:"error_message"`
}

func GetNewAddress(w http.ResponseWriter, r *http.Request) {
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

	checkAddress := CheckAddress(userID, params["currency"])
	if checkAddress {
		j, _ := json.Marshal("You already have " + params["currency"] + "address")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	body := NewAddressEndpoint(params["currency"], params["label"])
	if string(body) == "Error" {
		j, _ := json.Marshal("Block.io error")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}

	data := &NewAddress{
		Data: &Address{},
	}
	json.Unmarshal(body, data)
	log.Println(data.Data.Address)
	save := SaveAddress(userID, data.Data.Address, params["currency"], params["label"])

	if save {
		dataAdd := GetAddresses(userID)
		j, _ := json.Marshal(dataAdd)
		fmt.Println(string(j))
		w.WriteHeader(http.StatusOK)
		w.Write(j)
	}
}

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

	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetUserAddresses(w http.ResponseWriter, r *http.Request) {
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
	data := GetAddresses(id)

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
func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
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
func SendLocal(w http.ResponseWriter, r *http.Request) {
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
	checkAddress := CheckAddress(userID, params["currency"])
	if !checkAddress {
		j, _ := json.Marshal("Please create " + params["currency"] + " address")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	checkBalances := checkBalances(userID, params["currency"], params["amount"])
	if !checkBalances {
		j, _ := json.Marshal("Not enought " + params["currency"] + " balance")
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	destineToCheck := checkLabelDestine(params["to"], params["currency"])
	if !destineToCheck {
		j, _ := json.Marshal("Not label " + params["currency"] + " with name " + params["to"])
		w.WriteHeader(http.StatusBadRequest)
		w.Write(j)
		return
	}
	labelName := GetLabelFromID(params["currency"], userID)
	log.Printf(labelName)
	SendRequest := SendBalanceByLabels(params["currency"], params["amount"], labelName, params["to"])
	fmt.Println(string(SendRequest))

	data := &ResponseSendLocal{
		Data: &ResponseSendLocalData{},
	}
	json.Unmarshal(SendRequest, data)
	log.Println(data.Status)
	if data.Status == "success" {
		mapD := map[string]string{"status": "success", "txid": data.Data.TXID, "network": data.Data.Network, "sent": data.Data.ASent, "NetworkFee": data.Data.NetWorkFee}
		mapB, _ := json.Marshal(mapD)
		UpdateBalance(fmt.Sprintf("%v", claims["sub"]), params["currency"])
		w.WriteHeader(http.StatusOK)
		w.Write(mapB)
		return
	} else {
		mapD := map[string]string{"status": "error", "error": data.Data.ErrorMessage}
		mapB, _ := json.Marshal(mapD)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mapB)
		return
	}
	w.WriteHeader(http.StatusOK)
}
