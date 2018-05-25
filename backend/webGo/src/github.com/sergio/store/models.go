package store

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"../common"
)

func GetOrders() []GetOrdersStructure {
	var data []GetOrdersStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT id_order,amount,currency,price,create_at,currency_to FROM orders")
	if err != nil {
		log.Println(err.Error())
	}
	row := GetOrdersStructure{}
	for rows.Next() {
		var responseAmount string
		var responseCurrency string
		var responsePrice string
		var responseCreateAt string
		var currency_to string
		var id_order string
		err = rows.Scan(&id_order, &responseAmount, &responseCurrency, &responsePrice, &responseCreateAt, &currency_to)
		if err != nil {
			log.Println(err.Error())
		}
		row.ID = id_order
		row.Amount = responseAmount
		row.Currency = responseCurrency
		row.Price = responsePrice
		row.CreateAt = responseCreateAt
		row.CurrencyTo = currency_to
		data = append(data, row)
	}
	db.Close()
	return data
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
func GetUserOrders(id_account string) []GetOrdersStructure {
	var data []GetOrdersStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,currency,price,create_at FROM orders WHERE id_account=?", id_account)
	if err != nil {
		log.Println(err.Error())
	}
	row := GetOrdersStructure{}
	for rows.Next() {
		var responseAmount string
		var responseCurrency string
		var responsePrice string
		var responseCreateAt string
		err = rows.Scan(&responseAmount, &responseCurrency, &responsePrice, &responseCreateAt)
		if err != nil {
			log.Println(err.Error())
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
func SaveOrder(id_account string, amount string, currency string, price string, currency_to string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("INSERT INTO orders (id_account,amount,currency,price,create_at,currency_to) VALUES(?,?,?,?,?,?)")
	if err != nil {
		defer db.Close()
		return false
	}
	insForm.Exec(id_account, amount, currency, price, time.Now(), currency_to)
	defer db.Close()
	return true
}
func GetBalance(id_account string, currency string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT sum(amount) FROM orders WHERE id_account=? AND currency=?", id_account, currency)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		var amount string
		err = rows.Scan(&amount)
		if err != nil {
			log.Println(err.Error())
		}
		defer db.Close()
		return amount
	}
	defer db.Close()
	return ""
}
func SetPayment(amount string, id string, currency string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("UPDATE orders SET amount=? WHERE id_order=? AND currency=?")
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
func GetLabelFromId(id_account string, currency string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT label FROM addresses WHERE id_user=? AND currency=?", id_account, currency)
	if err != nil {
		log.Println(err.Error())
	}
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
func GetIdFromIdOrder(id_order string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT id_account FROM orders WHERE id_order=?", id_order)
	if err != nil {
		log.Println(err.Error())
	}
	for rows.Next() {
		var id_account string
		err = rows.Scan(&id_account)
		if err != nil {
			log.Println(err.Error())
		}
		defer db.Close()
		return id_account
	}
	defer db.Close()
	return ""
}
