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
	db := common.DbConn()
	fmt.Println(id)
	fmt.Println(address)
	fmt.Println(currency)
	insForm, err := db.Prepare("INSERT INTO addrs (id_user,address,label,currency,create_at) VALUES(?,?,?,?,?)")
	if err != nil {
		return false
	}
	insForm.Exec(id, address, label, currency, time.Now())
	defer db.Close()
	return true
}
func CheckAddress(id string, currency string) bool {
	db := common.DbConn()
	err := db.QueryRow("SELECT address,currency FROM addrs WHERE id_user=? AND currency=?", id, currency).Scan(&id, &currency)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that username.")
		return false
	case err != nil:
		log.Fatal(err)
		return false
	default:
		fmt.Printf("id is %s\n", id)
		return true
	}
	defer db.Close()
	return true
}
func GetGenericData(id_account string) []Data {
	var data []Data
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,currency FROM users_total WHERE id_user=?", id_account)
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
