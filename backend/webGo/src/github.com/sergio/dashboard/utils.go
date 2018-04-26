package dashboard

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func NewAddressEndpoint(currency string) []byte {
	res, err := http.Get(os.Getenv("BASE_API") + "/get_new_address/?api_key=" + os.Getenv(currency))
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
