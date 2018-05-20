package overview

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func UpdateBalancesApi(currency string, label string) []byte {
	fmt.Println("UpdateAPI")
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
