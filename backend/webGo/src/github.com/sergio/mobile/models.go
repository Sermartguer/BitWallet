package mobile

import (
	"log"

	"../common"
)

func GetUserData(id_account string) []MobileStructureLogin {
	var data []MobileStructureLogin
	db := common.DbConn()
	rows, err := db.Query("select username, acc_type, pin FROM accounts WHERE id=?", id_account)
	if err != nil {
		log.Println(err.Error())
	}
	row := MobileStructureLogin{}
	for rows.Next() {
		var username string
		var account string
		var pin string
		err = rows.Scan(&username, &account, &pin)
		if err != nil {
			log.Println(err.Error())
		}
		row.Username = username
		row.Account = account
		row.Pin = pin
		data = append(data, row)
	}
	db.Close()
	return data
}
func GetUserBalance(id_account string) []MobileBalances {
	var data []MobileBalances
	db := common.DbConn()
	rows, err := db.Query("select currency, amount FROM addresses WHERE id_user=?", id_account)
	if err != nil {
		log.Println(err.Error())
	}
	row := MobileBalances{}
	for rows.Next() {
		var currency string
		var balance string
		err = rows.Scan(&currency, &balance)
		if err != nil {
			log.Println(err.Error())
		}
		row.Currency = currency
		row.Balance = balance
		data = append(data, row)
	}
	db.Close()
	return data
}
