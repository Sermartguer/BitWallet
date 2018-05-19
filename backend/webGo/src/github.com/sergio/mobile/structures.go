package mobile

type MobileStructureLogin struct {
	Username string `json:"username"`
	Account  string `json:"account"`
	Pin      string `json:"pin"`
}
type MobileBalances struct {
	Currency string `json:"currency"`
	Balance  string `json:"balance"`
}
