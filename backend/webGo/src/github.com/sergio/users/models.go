package users

import (
	"database/sql"
	"fmt"
	"log"

	"../common"
)

func SaveUser(user_data UserModelValidator) bool {
	db := common.DbConn()

	insForm, err := db.Prepare("INSERT INTO users (id, username, email,password,acc_type,update_at, create_at,active) VALUES(?,?,?,?,?,?,?,?)")
	if err != nil {
		return false
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
	err := db.QueryRow("SELECT username,email FROM users WHERE username=? OR email=? ", username, email).Scan(&username, &email)
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
	err := db.QueryRow("SELECT username FROM users WHERE username=?", username).Scan(&username)
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
	err := db.QueryRow("SELECT password FROM users WHERE username=?", username).Scan(&username)
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
	err := db.QueryRow("SELECT email FROM users WHERE username=?", username).Scan(&username)
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
