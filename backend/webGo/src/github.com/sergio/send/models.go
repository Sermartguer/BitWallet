package send

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

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
		fmt.Printf("email is %s\n", username)
	}
	defer db.Close()
	return username
}
func GetAddresses(id string) []GetAddressesStructure {
	var data []GetAddressesStructure
	db := common.DbConn()
	rows, err := db.Query("CALL get_accounts(?)", id)
	if err != nil {
		panic(err.Error())
	}
	row := GetAddressesStructure{}
	for rows.Next() {
		var currency string
		var address string
		var label string
		err = rows.Scan(&address, &currency, &label)
		if err != nil {
			panic(err.Error())
		}
		row.Currency = currency
		row.Address = address
		row.Label = label
		data = append(data, row)
	}
	db.Close()
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
		log.Println(activo)
		if activo == "0" {
			return false
		} else {
			return true
		}
	}
	return true
}
func CheckPin(id string, pin string) bool {
	db := common.DbConn()
	rows, err := db.Query("SELECT pin FROM accounts WHERE id=?", id)
	for rows.Next() {
		var pinDB string
		err = rows.Scan(&pinDB)
		if err != nil {
			log.Println(err.Error())
		}
		if pinDB == pin {
			return true
		} else {
			return false
		}
	}
	return false
}
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
func SaveAddress(id string, address string, currency string, label string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("UPDATE addresses SET address=?,label=?,active=1 WHERE id_user=? AND currency=?")
	if err != nil {
		panic(err)
	}
	_, err = insForm.Exec(address, label, id, currency)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()
	return true

}
func checkBalances(id string, currency string, amount string) bool {
	db := common.DbConn()
	rows, err := db.Query("SELECT amount FROM addresses WHERE id_user=? AND currency=?", id, currency)
	for rows.Next() {
		var label string
		err = rows.Scan(&label)
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(label)
		fmt.Println(amount)
		labelValue, _ := strconv.ParseFloat(label, 64)
		amountValue, _ := strconv.ParseFloat(amount, 64)
		if labelValue > amountValue {
			return true
		} else {
			return false
		}
	}
	return false
}
func checkLabelDestine(destine string, currency string) bool {
	db := common.DbConn()
	rows, err := db.Query("SELECT active FROM addresses WHERE label=? AND currency=?", destine, currency)
	for rows.Next() {
		var label string
		err = rows.Scan(&label)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println()
		if label == "1" {
			return true
		} else {
			return false
		}
	}
	return false
}
func GetLabelFromID(currency string, id string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT label FROM addresses WHERE id_user=? AND currency=?", id, currency)
	for rows.Next() {
		var label string
		err = rows.Scan(&label)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(label)
		return label
	}
	return ""
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
func SaveTransaction(id string, to string, hash_id string, amount string, currency string, trans_type string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("INSERT INTO transactions (id_account,send_to,hash_id,amount,currency,trans_type,trans_time) VALUES(?,?,?,?,?,?,?)")
	if err != nil {
		return false
	}
	insForm.Exec(id, to, hash_id, "-"+amount, currency, trans_type, time.Now())
	defer db.Close()
	return true
}
func UpdateBalanceTo(amount string, id string, currency string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("UPDATE addresses SET amount=? WHERE id_user=? AND currency=?")
	if err != nil {
		panic(err)
	}
	_, err = insForm.Exec(amount, id, currency)

	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()
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
		return label
	}
	return ""
}
