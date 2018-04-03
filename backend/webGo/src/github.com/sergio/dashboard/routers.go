package dashboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"../common"
)

type NewAddress struct {
	Status string   `json:"status"`
	Data   *Address `json:"data"`
}
type Address struct {
	Address string `json:"address"`
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

	body := NewAddressEndpoint(params["currency"])
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
	s2, _ := json.Marshal(data)

	save := SaveAddress(userID, data.Data.Address, params["currency"])
	if save {
		w.WriteHeader(http.StatusOK)
		w.Write(s2)
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
