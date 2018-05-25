package mobile

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

func SendBalanceByAddress(currency string, amount string, from string, to string) []byte {
	newAmount := 0.0
	s, err := strconv.ParseFloat(amount, 32)
	if err == nil {
		newAmount = s
	}
	a := fmt.Sprintf("%f", newAmount)
	res, err := http.Get(os.Getenv("BASE_API") + "/withdraw_from_labels/?api_key=" + os.Getenv(currency) + "&amounts=" + a + "&from_labels=" + from + "&to_labels=" + to + "&pin=" + os.Getenv("BlockioPin"))
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
