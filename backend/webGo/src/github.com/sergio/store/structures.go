package store

type GetOrdersStructure struct {
	ID         string `json:"id"`
	Amount     string `json:"amount"`
	Currency   string `json:"currency"`
	Price      string `json:"price"`
	CreateAt   string `json:"create_at"`
	CurrencyTo string `json:"currency_to"`
}
