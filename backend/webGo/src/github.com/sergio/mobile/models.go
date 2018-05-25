package mobile

import (
	"log"
	"time"

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
	defer db.Close()
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
	defer db.Close()
	return data
}
func CreateOrderMobile(currency string, amount string, mobileID string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("INSERT INTO mobileOrders (id_account,currency,amount,time) VALUES(?,?,?,?)")
	if err != nil {
		defer db.Close()
		return false
	}
	insForm.Exec(mobileID, currency, amount, time.Now())
	defer db.Close()
	return true
}
func GetLabelFromMobileId(id string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT label FROM mobileOrders INNER JOIN addresses ON mobileOrders.id_account = addresses.id_user AND mobileOrders.currency = addresses.currency WHERE id_account = ?", id)
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
func GetCurrencyFromMobileId(id string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT mobileOrders.currency FROM mobileOrders INNER JOIN addresses ON mobileOrders.id_account = addresses.id_user AND mobileOrders.currency = addresses.currency WHERE id_account = ?", id)
	for rows.Next() {
		var currency string
		err = rows.Scan(&currency)
		if err != nil {
			log.Println(err.Error())
		}
		defer db.Close()
		return currency
	}
	defer db.Close()
	return ""
}
func GetLabelFromParamId(currency string, id string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT label FROM addresses WHERE currency=? AND id_user=?", currency, id)
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
func GetAmountId(currency string, id string) string {
	db := common.DbConn()
	rows, err := db.Query("SELECT amount FROM mobileOrders WHERE currency=? AND id_account=?", currency, id)
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
func DeleteOrder(id string) bool {
	db := common.DbConn()
	_, err := db.Query("DELETE FROM mobileOrders WHERE id_account=?", id)
	if err != nil {
		defer db.Close()
		return false
	} else {
		defer db.Close()
		return true
	}
}
