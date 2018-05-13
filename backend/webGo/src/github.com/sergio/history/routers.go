package history

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"../common"
)

func GetLoginHistory(w http.ResponseWriter, r *http.Request) {
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
	data := GetLoginLog(id)
	j, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetActionsHistory(w http.ResponseWriter, r *http.Request) {
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
	data := GetActionsLog(id)
	j, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
func GetOrderHistory(w http.ResponseWriter, r *http.Request) {
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
	data := GetOrderLog(id)
	j, _ := json.Marshal(data)
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
