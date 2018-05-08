package dashboard

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func NewAddressEndpoint(currency string, label string) []byte {
	res, err := http.Get(os.Getenv("BASE_API") + "/get_new_address/?api_key=" + os.Getenv(currency) + "&label=" + label)
	if err != nil {
		return []byte("Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("Error")
	} else {
		return body
	}
}
func GetCurrencyPrice(APIkey string) []byte {
	key := APIkey
	log.Printf(key)
	res, err := http.Get(os.Getenv("BASE_API") + "/get_current_price/?api_key=" + key)
	if err != nil {
		return []byte("Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("Error")
	} else {
		return body
	}
}
func UpdateBalancesApi(currency string, label string) []byte {
	res, err := http.Get(os.Getenv("BASE_API") + "/get_address_by/?api_key=" + os.Getenv(currency) + "&label=" + label)
	if err != nil {
		return []byte("Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("Error")
	} else {
		return body
	}
}
func SendBalanceByLabels(currency string, amount string, from string, to string) []byte {
	res, err := http.Get(os.Getenv("BASE_API") + "/withdraw_from_labels/?api_key=" + os.Getenv(currency) + "&amounts=" + amount + "&from_labels=" + from + "&to_labels=" + to + "&pin=" + os.Getenv("BlockioPin"))
	if err != nil {
		return []byte("Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("Error")
	} else {
		return body
	}
}
func SendBalanceByAddress(currency string, amount string, from string, to string) []byte {
	res, err := http.Get(os.Getenv("BASE_API") + "/withdraw_from_labels/?api_key=" + os.Getenv(currency) + "&amounts=" + amount + "&from_labels=" + from + "&to_addresses=" + to + "&pin=" + os.Getenv("BlockioPin"))
	if err != nil {
		return []byte("Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("Error")
	} else {
		return body
	}
}
func GetFee(currency string, amount string, address string) []byte {
	res, err := http.Get(os.Getenv("BASE_API") + "/get_network_fee_estimate/?api_key=" + os.Getenv(currency) + "&amounts=" + amount + "&to_addresses=" + address)
	if err != nil {
		return []byte("Error")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte("Error")
	} else {
		return body
	}
}
