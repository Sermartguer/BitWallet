package history

import (
	"database/sql"
	"log"

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

//GetLoginLog used to get login log from user
func GetLoginLog(id_account string) []LoginStructure {
	var data []LoginStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT ip,time,success FROM activity_login WHERE id_account=?", id_account)
	if err != nil {
		panic(err.Error())
	}
	row := LoginStructure{}
	for rows.Next() {
		var ip string
		var time string
		var success string
		err = rows.Scan(&ip, &time, &success)
		if err != nil {
			panic(err.Error())
		}
		row.IP = ip
		row.Time = time
		row.Success = success
		data = append(data, row)
	}
	db.Close()
	return data
}

//GetActionsLog used to get login log from user
func GetActionsLog(id_account string) []ActionStructure {
	var data []ActionStructure
	db := common.DbConn()
	rows, err := db.Query("SELECT amount,address_local,currency,time,action FROM activity_actions WHERE id_account=?", id_account)
	if err != nil {
		panic(err.Error())
	}
	row := ActionStructure{}
	for rows.Next() {
		var amount string
		var address string
		var currency string
		var time string
		var action string
		err = rows.Scan(&amount, &address, &currency, &time, &action)
		if err != nil {
			panic(err.Error())
		}
		row.Amount = amount
		row.Address = address
		row.Currency = currency
		row.Time = time
		row.Action = action
		data = append(data, row)
	}
	db.Close()
	return data
}

//GetOrderLog used to get login log from user
func GetOrderLog(id_account string) []Ordertructure {
	var data []Ordertructure
	db := common.DbConn()
	rows, err := db.Query("SELECT action,amount,price,currency,time FROM activity_orders WHERE id_account=?", id_account)
	if err != nil {
		panic(err.Error())
	}
	row := Ordertructure{}
	for rows.Next() {
		var amount string
		var currency string
		var time string
		var price string
		var action string
		err = rows.Scan(&action, &amount, &price, &currency, &time)
		if err != nil {
			panic(err.Error())
		}
		row.Amount = amount
		row.Price = price
		row.Currency = currency
		row.Time = time
		row.Action = action
		data = append(data, row)
	}
	db.Close()
	return data
}
