package common

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

//DbConn used to connect server to mysql database
func DbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "BitWallet"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		log.Println(err.Error())
	}
	return db
}
