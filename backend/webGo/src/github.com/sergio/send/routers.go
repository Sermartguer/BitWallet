package send

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"../common"
)

func GetUserAddresses(w http.ResponseWriter, r *http.Request) {
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
	data := GetAddresses(id)

	j, _ := json.Marshal(data)
	fmt.Println(string(j))

	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetNewAddress(w http.ResponseWriter, r *http.Request) {
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

	checkAddress := CheckAddress(userID, params["currency"])
	if checkAddress {
		common.StatusBadError(w, r, "You already have "+params["currency"]+"address")
		return
	}
	body := NewAddressEndpoint(params["currency"], params["label"])
	if string(body) == "Error" {
		common.StatusBadError(w, r, "Block.io error")
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
func SendLocal(w http.ResponseWriter, r *http.Request) {
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
	pinChecked := CheckPin(userID, params["pin"])
	if !pinChecked {
		common.StatusBadError(w, r, "Pin error")
		return
	}
	checkAddress := CheckAddress(userID, params["currency"])
	if !checkAddress {
		common.StatusBadError(w, r, "Please create "+params["currency"]+" address")
		return
	}
	checkBalances := checkBalances(userID, params["currency"], params["amount"])
	if !checkBalances {
		common.StatusBadError(w, r, "Not enought "+params["currency"]+" balance")
		return
	}
	destineToCheck := checkLabelDestine(params["to"], params["currency"])
	if !destineToCheck {
		common.StatusBadError(w, r, "Not label "+params["currency"]+" with name "+params["to"])
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
		save := SaveTransaction(userID, params["to"], data.Data.TXID, params["amount"], params["currency"], "local")
		if save {
			mapD := map[string]string{"status": "success", "txid": data.Data.TXID, "network": data.Data.Network, "sent": data.Data.ASent, "NetworkFee": data.Data.NetWorkFee}
			mapB, _ := json.Marshal(mapD)
			UpdateBalance(username, params["currency"])
			w.WriteHeader(http.StatusOK)
			w.Write(mapB)
			return
		} else {
			common.StatusBadError(w, r, "Server error")
			return
		}
	} else {
		mapB, _ := json.Marshal(data.Data.ErrorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mapB)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func SendExternal(w http.ResponseWriter, r *http.Request) {
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
	pinChecked := CheckPin(userID, params["pin"])
	if !pinChecked {
		common.StatusBadError(w, r, "Pin error")
		return
	}
	checkAddress := CheckAddress(userID, params["currency"])
	if !checkAddress {
		common.StatusBadError(w, r, "Please create "+params["currency"]+" address")
		return
	}
	checkBalances := checkBalances(userID, params["currency"], params["amount"])
	if !checkBalances {
		common.StatusBadError(w, r, "Not enought "+params["currency"]+" balance")
		return
	}

	labelName := GetLabelFromID(params["currency"], userID)
	log.Printf(labelName)
	SendRequest := SendBalanceByAddress(params["currency"], params["amount"], labelName, params["to"])
	fmt.Println(string(SendRequest))

	data := &ResponseSendLocal{
		Data: &ResponseSendLocalData{},
	}
	json.Unmarshal(SendRequest, data)
	log.Println(data.Status)
	if data.Status == "success" {
		save := SaveTransaction(userID, params["to"], data.Data.TXID, params["amount"], params["currency"], "external")
		if save {
			mapD := map[string]string{"status": "success", "txid": data.Data.TXID, "network": data.Data.Network, "sent": data.Data.ASent, "NetworkFee": data.Data.NetWorkFee}
			mapB, _ := json.Marshal(mapD)
			UpdateBalance(username, params["currency"])
			w.WriteHeader(http.StatusOK)
			w.Write(mapB)
			return
		} else {
			common.StatusBadError(w, r, "Server error")
			return
		}
	} else {
		mapB, _ := json.Marshal(data.Data.ErrorMessage)
		w.WriteHeader(http.StatusBadRequest)
		w.Write(mapB)
		return
	}
	w.WriteHeader(http.StatusOK)
}
func GetNetworkFee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)

	username, errorToken := common.GetUsernameByToken(params["token"])
	if errorToken {
		common.StatusBadError(w, r, "Error in token check")
		return
	}
	GetNetworkFee := GetFee(params["currency"], params["amount"], params["address"])

	data := &ResponseGetNetworkFee{
		Data: &ResponseGetNetworkFeeData{},
	}
	json.Unmarshal(GetNetworkFee, data)
	if data.Status == "success" {
		mapD := map[string]string{"status": "success", "network_fee": data.Data.EstimatedFee}
		mapB, _ := json.Marshal(mapD)
		UpdateBalance(username, params["currency"])
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
		UpdateBalanceTo(data.Data.Balance, userID, currency)
	}
}
