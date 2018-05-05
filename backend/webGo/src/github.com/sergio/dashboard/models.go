package dashboard

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"../common"
)

type Data struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
type GetAddressesStructure struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Label    string `json:"label"`
}
type GetOrdersStructure struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Price    string `json:"price"`
	CreateAt string `json:"create_at"`
}
type TransactionsStructure struct {
	SendTo    string `json:"send_to"`
	HashId    string `json:"hash_id"`
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	TransTime string `json:"trans_time"`
}
type AddressStruct struct {
	Activo string `json:"activo"`
}

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

func SaveAddress(id string, address string, currency string, label string) bool {
	log.Println("SaveAddress")
	log.Println(id)
	log.Println(address)
	log.Println(currency)
	log.Println(label)
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
func CheckAddress(id string, currency string) bool {
	log.Println("CheckAddress")
	log.Println(id)
	log.Println(currency)
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
func GetGenericData(id_account string) []Data {
	var data []Data
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,currency FROM addresses WHERE id_user=?", id_account)
	if err != nil {
		panic(err.Error())
	}
	row := Data{}
	for rows.Next() {
		var responseAmount string
		var responseCurrency string
		err = rows.Scan(&responseAmount, &responseCurrency)
		if err != nil {
			panic(err.Error())
		}
		row.Amount = responseAmount
		row.Currency = responseCurrency
		data = append(data, row)
	}
	db.Close()
	return data
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
func SaveOrder(id_account string, amount string, currency string, price string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("INSERT INTO orders (id_account,amount,currency,price,create_at) VALUES(?,?,?,?,?)")
	if err != nil {
		return false
	}
	insForm.Exec(id_account, amount, currency, price, time.Now())
	defer db.Close()
	return true
}
func GetOrders() []GetOrdersStructure {
	var data []GetOrdersStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,currency,price,create_at FROM orders")
	if err != nil {
		panic(err.Error())
	}
	row := GetOrdersStructure{}
	for rows.Next() {
		var responseAmount string
		var responseCurrency string
		var responsePrice string
		var responseCreateAt string
		err = rows.Scan(&responseAmount, &responseCurrency, &responsePrice, &responseCreateAt)
		if err != nil {
			panic(err.Error())
		}
		row.Amount = responseAmount
		row.Currency = responseCurrency
		row.Price = responsePrice
		row.CreateAt = responseCreateAt
		data = append(data, row)
	}
	db.Close()
	return data
}
func GetUserOrders(id_account string) []GetOrdersStructure {
	var data []GetOrdersStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,currency,price,create_at FROM orders WHERE id_account=?", id_account)
	if err != nil {
		panic(err.Error())
	}
	row := GetOrdersStructure{}
	for rows.Next() {
		var responseAmount string
		var responseCurrency string
		var responsePrice string
		var responseCreateAt string
		err = rows.Scan(&responseAmount, &responseCurrency, &responsePrice, &responseCreateAt)
		if err != nil {
			panic(err.Error())
		}
		row.Amount = responseAmount
		row.Currency = responseCurrency
		row.Price = responsePrice
		row.CreateAt = responseCreateAt
		data = append(data, row)
	}
	db.Close()
	return data
}
func UserTransactions(id_account string) []TransactionsStructure {
	var data []TransactionsStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT send_to,hash_id,amount,currency,trans_time FROM transactions WHERE id_account=?", id_account)
	if err != nil {
		panic(err.Error())
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
			panic(err.Error())
		}
		row.SendTo = responseSendTo
		row.HashId = responseHashId
		row.Amount = responseAmount
		row.Currency = responseCurrency
		row.TransTime = responseTransTime
		data = append(data, row)
	}
	db.Close()
	return data
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
func UpdateBalanceTo(amount string, id string, currency string) bool {
	log.Println("Update Balance")
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
