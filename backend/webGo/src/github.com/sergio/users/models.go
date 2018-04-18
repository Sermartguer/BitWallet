package users

import (
	"database/sql"
	"fmt"
	"log"

	"../common"
)

func SaveUser(user_data UserModelValidator) bool {
	db := common.DbConn()

	insForm, err := db.Prepare("INSERT INTO accounts (id, username, email,password,acc_type,update_at, create_at,active) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return false
		log.Fatal(err)
	}
	insForm.Exec(user_data.ID, user_data.Username, user_data.Email, user_data.Password, user_data.AccountType, user_data.CreatedAt, user_data.CreatedAt, false)
	defer db.Close()
	return true
}

func CheckUser(user_data UserModelValidator) bool {
	db := common.DbConn()
	//Check if username exist on DB
	username := user_data.Username
	email := user_data.Email
	err := db.QueryRow("SELECT username,email FROM accounts WHERE username=? OR email=? ", username, email).Scan(&username, &email)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that username.")
		return false
	case err != nil:
		log.Fatal(err)
	default:
		return true
		fmt.Printf("Username is %s\n", username)
	}
	defer db.Close()
	return true
}
func SearchUser(username string) bool {
	db := common.DbConn()
	err := db.QueryRow("SELECT username FROM accounts WHERE username=?", username).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that username.")
		return false
	case err != nil:
		log.Fatal(err)
		return false
	default:
		return true
		fmt.Printf("Username is %s\n", username)
	}
	defer db.Close()
	return false
}
func GetPassword(username string) string {
	db := common.DbConn()
	err := db.QueryRow("SELECT password FROM accounts WHERE username=?", username).Scan(&username)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that username.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("password is %s\n", username)
	}
	defer db.Close()
	return username
}
func GetEmail(username string) string {
	db := common.DbConn()
	err := db.QueryRow("SELECT email FROM accounts WHERE username=?", username).Scan(&username)
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
func Verify(id string) bool {
	db := common.DbConn()

	insForm, err := db.Prepare("UPDATE accounts SET active=true WHERE id=?")
	if err != nil {
		panic(err)
	}
	_, err = insForm.Exec(id)

	if err != nil {
		log.Fatal(err)
		return false

	}
	defer db.Close()
	return true
}
