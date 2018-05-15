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
	rows, err := db.Query("SELECT amount,currency,price,create_at FROM orders")
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
