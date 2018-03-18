package users

import (
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
