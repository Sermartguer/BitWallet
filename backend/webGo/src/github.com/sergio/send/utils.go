package send

import (
	"io/ioutil"
	"net/http"
	"os"
)

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
