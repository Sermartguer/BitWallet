package users

import (
	"database/sql"
	"fmt"
	"log"

	"../common"
)

type AccountStruct struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"firstname"`
	Surname    string `json:"surname"`
	AcountType string `json:"acc_type"`
}

func SaveUser(user_data UserModelValidator, mobile string) bool {
	db := common.DbConn()

	insForm, err := db.Prepare("INSERT INTO accounts (id, username, email,password,acc_type,update_at, create_at, active, mobile_hash) VALUES(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		return false
		log.Fatal(err)
	}
	insForm.Exec(user_data.ID, user_data.Username, user_data.Email, user_data.Password, user_data.AccountType, user_data.CreatedAt, user_data.CreatedAt, false, mobile)
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
func Verify(id string, pin string) bool {
	db := common.DbConn()

	insForm, err := db.Prepare("UPDATE accounts SET active=true, pin=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	_, err = insForm.Exec(pin, id)

	if err != nil {
		log.Fatal(err)
		return false

	}
	defer db.Close()
	return true
}
func Update(firstname string, surname string, id string) bool {
	log.Println(firstname)
	db := common.DbConn()
	insForm, err := db.Prepare("UPDATE accounts SET firstname=?, surname=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	_, err = insForm.Exec(firstname, surname, id)

	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()
	return true
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
func GetIdByEmail(email string) string {
	db := common.DbConn()
	err := db.QueryRow("SELECT id FROM accounts WHERE email=?", email).Scan(&email)
	switch {
	case err == sql.ErrNoRows:
		log.Printf("No user with that username.")
	case err != nil:
		log.Fatal(err)
	default:
		fmt.Printf("email is %s\n", email)
	}
	defer db.Close()
	return email
}
func GetAccount(id string) []AccountStruct {
	var data []AccountStruct
	db := common.DbConn()
	rows, err := db.Query("SELECT username,email,firstname,surname,acc_type,mobile_hash FROM accounts WHERE id=?", id)
	if err != nil {
		log.Printf("Err1")
		log.Printf(err.Error())
	}
	row := AccountStruct{}
	for rows.Next() {
		var responseUsername string
		var responseEmail string
		var responseFirstName string
		var responseSurname string
		var responseAcountType string
		err = rows.Scan(&responseUsername, &responseEmail, &responseFirstName, &responseSurname, &responseAcountType)
		if err != nil {
			log.Printf("Err2")
			log.Printf(err.Error())
		}
		row.Username = responseUsername
		row.Email = responseEmail
		row.FirstName = responseFirstName
		row.Surname = responseSurname
		row.AcountType = responseAcountType
		data = append(data, row)
	}
	db.Close()
	return data
}
func NewAccountPassword(password string, id string) bool {
	db := common.DbConn()
	insForm, err := db.Prepare("UPDATE accounts SET password=? WHERE id=?")
	if err != nil {
		panic(err)
	}
	_, err = insForm.Exec(password, id)

	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()
	return true
}
