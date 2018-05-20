package store

import (
	"io/ioutil"
	"net/http"
	"os"
)

func SendBalanceByAddress(currency string, amount string, from string, to string) []byte {
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
