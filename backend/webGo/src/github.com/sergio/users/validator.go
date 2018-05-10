package users

import (
	"encoding/json"
	"log"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/rs/xid"
	"golang.org/x/crypto/bcrypt"
)

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
	user_data.Password2 = params["password2"]
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
	user_data.Password, _ = HashPassword(user_data.Password)
	user_data.Error = true
	return user_data
}
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
