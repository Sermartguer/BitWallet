package history

type LoginStructure struct {
	IP      string `json:"ip"`
	Time    string `json:"time"`
	Success string `json:"success"`
}
type ActionStructure struct {
	Amount   string `json:"amount"`
	Address  string `json:"address_local"`
	Currency string `json:"currency"`
	Time     string `json:"time"`
	Action   string `json:"action"`
}
type Ordertructure struct {
	Action   string `json:"action"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
	Price    string `json:"price"`
	Time     string `json:"time"`
}
