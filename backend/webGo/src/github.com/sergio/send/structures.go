package send

type ResponseSendLocal struct {
	Status string                 `json:"status"`
	Data   *ResponseSendLocalData `json:"data"`
}
type ResponseSendLocalData struct {
	Network      string `json:"network"`
	TXID         string `json:"txid"`
	AWithdraw    string `json:"amount_withdrawn"`
	ASent        string `json:"amount_sent"`
	NetWorkFee   string `json:"network_fee"`
	BlockIOFee   string `json:"blockio_fee"`
	ErrorMessage string `json:"error_message"`
}
type ResponseGetBalance struct {
	Status string   `json:"status"`
	Data   *Balance `json:"data"`
}
type Balance struct {
	Balance string `json:"available_balance"`
}
type ResponseGetNetworkFee struct {
	Status string                     `json:"status"`
	Data   *ResponseGetNetworkFeeData `json:"data"`
}
type ResponseGetNetworkFeeData struct {
	Network      string `json:"network"`
	EstimatedFee string `json:"estimated_network_fee"`
	ErrorMessage string `json:"error_message"`
}
type NewAddress struct {
	Status string   `json:"status"`
	Data   *Address `json:"data"`
}
type Address struct {
	Address string `json:"address"`
}
type GetAddressesStructure struct {
	Address  string `json:"address"`
	Currency string `json:"currency"`
	Label    string `json:"label"`
}
