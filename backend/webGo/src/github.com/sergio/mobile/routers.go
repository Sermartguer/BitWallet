package mobile

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"../common"
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

func MobilePay(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)
	data := GetLabelFromMobileId(params["mobileID"])
	//Usuari ha generat el qr
	log.Println("a")
	if data != "" {
		dataCurrency := GetCurrencyFromMobileId(params["mobileID"])
		if dataCurrency != "" {
			//Usuari fa scan
			log.Println("a")
			dataQR := GetLabelFromParamId(dataCurrency, params["payID"])
			if dataQR != "" {
				log.Println("a")
				amount := GetAmountId(dataCurrency, params["mobileID"])
				log.Println(amount)
				status := SendBalanceByAddress(dataCurrency, amount, dataQR, data)
				log.Println(string(status))
				delete := DeleteOrder(params["mobileID"])
				if delete {
					w.WriteHeader(http.StatusOK)
					return
				} else {
					common.StatusBadError(w, r, "Error on delete")
					return
				}
			} else {
				common.StatusBadError(w, r, "Error in order")
				return
			}
		} else {
			common.StatusBadError(w, r, "Error no currency found")
			return
		}
	} else {
		common.StatusBadError(w, r, "Error no label found")
		return
	}
}
func GenerateMobileOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	dat, _ := ioutil.ReadAll(r.Body)
	var params map[string]string
	json.Unmarshal(dat, &params)
	data := CreateOrderMobile(params["currency"], params["amount"], params["mobileID"])
	if data {
		w.WriteHeader(http.StatusOK)
		return
	} else {
		common.StatusBadError(w, r, "Error to create order")
		return
	}
}
