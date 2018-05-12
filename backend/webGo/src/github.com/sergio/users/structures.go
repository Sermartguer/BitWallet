package users

type AccountStruct struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	FirstName  string `json:"firstname"`
	Surname    string `json:"surname"`
	AcountType string `json:"acc_type"`
	MobileHash string `json:"mobile_hash"`
}

type Error struct {
	TextError string `json:"error"`
}
type UserModelValidator struct {
	ID          string `json:"id" valid:"-"`
	Username    string `json:"username" valid:"required~Username is blank"`
	Email       string `json:"email" valid:"email"`
	Password    string `json:"password" valid:"length(5|10)"`
	Password2   string `json:"password2" valid:"length(5|10)"`
	CreatedAt   string `json:"create_at" valid:"-"`
	AccountType string `json:"acc_type" valid:"optional"`
	Error       bool   `json:"error" valid:"optional"`
	TextError   string `json:"text_error" valid:"optional"`
}
