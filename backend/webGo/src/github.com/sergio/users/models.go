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
