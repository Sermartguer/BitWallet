package dashboard

import (
	"database/sql"
	"fmt"
	"log"
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

func SaveAddress(id string, address string, currency string) bool {
	db := common.DbConn()
	fmt.Println(id)
	fmt.Println(address)
	fmt.Println(currency)
	insForm, err := db.Prepare("INSERT INTO addrs (id_user,address,currency,create_at) VALUES(?,?,?,?)")
	if err != nil {
		return false
	}
	insForm.Exec(id, address, currency, time.Now())
	defer db.Close()
	return true
}
