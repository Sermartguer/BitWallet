package users

import (
	"encoding/json"
	"log"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
)

type UserModelValidator struct {
	ID          string `json:"id" valid:"-"`
	Username    string `json:"username" valid:"required~Username is blank"`
	Email       string `json:"email" valid:"email"`
	Password    string `json:"password" valid:"length(5|10)"`
	CreatedAt   string `json:"create_at" valid:"-"`
	AccountType string `json:"acc_type" valid:"optional"`
	Error       bool   `json:"error" valid:"optional"`
	TextError   string `json:"text_error" valid:"optional"`
}

func ValidateParams(data []byte) UserModelValidator {
	user_data := UserModelValidator{}
	// Unmarshall this into a map
	var params map[string]string
	json.Unmarshal(data, &params)

	var id = xid.New()
	user_data.ID = id.String()
	user_data.Username = params["username"]
	user_data.Email = params["email"]
	user_data.Password = params["password"]
	user_data.AccountType = params["acc_type"]
	user_data.CreatedAt = time.Now().Format("Mon Jan _2 15:04:05 2006")

	result, err := govalidator.ValidateStruct(user_data)
	if err != nil {
		println("error: " + err.Error())
		user_data.TextError = err.Error()
		log.Println(result)
		user_data.Error = result
		return user_data

	}
	user_data.Error = true
	return user_data
}
