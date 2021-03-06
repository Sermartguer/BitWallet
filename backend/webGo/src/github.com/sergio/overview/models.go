package overview

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"../common"
)

func GetIdByUsername(username string) string {
	db := common.DbConn()
	err := db.QueryRow("SELECT id FROM accounts WHERE username=?", username).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that username.")
	case err != nil:
		log.Fatal(err)
	default:
	}
	defer db.Close()
	return username
}
func GetGenericData(id_account string) []Data {
	var data []Data
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,currency FROM addresses WHERE id_user=?", id_account)
	if err != nil {
		log.Println(err.Error())
	}
	row := Data{}
	for rows.Next() {
		var responseAmount string
		var responseCurrency string
		err = rows.Scan(&responseAmount, &responseCurrency)
		if err != nil {
			log.Println(err.Error())
		}
		row.Amount = responseAmount
		row.Currency = responseCurrency
		data = append(data, row)
	}
	defer db.Close()
	return data
}
func CheckAddress(id string, currency string) bool {
	db := common.DbConn()
	rows, err := db.Query("SELECT active FROM addresses WHERE id_user=? AND currency=?", id, currency)
	for rows.Next() {
		var activo string
		err = rows.Scan(&activo)
		if err != nil {
			log.Println(err.Error())
		}
		if activo == "0" {
			defer db.Close()
			return false
		} else {
			defer db.Close()
			return true
		}
	}
	return true
}
func GetLabelByID(id string, currency string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT label FROM addresses WHERE id_user=? AND currency=?", id, currency)
	for rows.Next() {
		var label string
		err = rows.Scan(&label)
		if err != nil {
			log.Println(err.Error())
		}
		defer db.Close()
		return label
	}
	defer db.Close()
	return ""
}
func UpdateBalanceTo(amount string, id string, currency string) bool {
	fmt.Println("a")
	fmt.Println(amount)
	db := common.DbConn()
	insForm, err := db.Prepare("UPDATE addresses SET amount=? WHERE id_user=? AND currency=?")
	if err != nil {
		log.Println(err.Error())
	}
	_, err = insForm.Exec(amount, id, currency)

	if err != nil {
		log.Fatal(err)
		defer db.Close()
		return false
	}
	defer db.Close()
	return true
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
func UserTransactions(id_account string) []TransactionsStructure {
	var data []TransactionsStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT send_to,hash_id,amount,currency,trans_time FROM transactions WHERE id_account=?", id_account)
	if err != nil {
		log.Println(err.Error())
	}
	row := TransactionsStructure{}
	for rows.Next() {
		var responseSendTo string
		var responseHashId string
		var responseAmount string
		var responseCurrency string
		var responseTransTime string
		err = rows.Scan(&responseSendTo, &responseHashId, &responseAmount, &responseCurrency, &responseTransTime)
		if err != nil {
			log.Println(err.Error())
		}
		row.SendTo = responseSendTo
		row.HashId = responseHashId
		row.Amount = responseAmount
		row.Currency = responseCurrency
		row.TransTime = responseTransTime
		data = append(data, row)
	}
	defer db.Close()
	return data
}
