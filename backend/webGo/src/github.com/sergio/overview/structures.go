package overview

type Data struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}
type TransactionsStructure struct {
	SendTo    string `json:"send_to"`
	HashId    string `json:"hash_id"`
	Amount    string `json:"amount"`
	Currency  string `json:"currency"`
	TransTime string `json:"trans_time"`
}
type ResponseGetBalance struct {
	Status string   `json:"status"`
	Data   *Balance `json:"data"`
}
type Balance struct {
	Balance string `json:"available_balance"`
}
