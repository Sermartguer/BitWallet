package store

type GetOrdersStructure struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
	Price    string `json:"price"`
	CreateAt string `json:"create_at"`
}
