package coinpayments

// CreatePaymentResponse is the api response after creating a payment
type CreatePaymentResponse struct {
	Error  string              `json:"error"`
	Result CreatePaymentResult `json:"result"`
}

// CreatePaymentResult is the result of creating a payment
type CreatePaymentResult struct {
	Amount         string `json:"amount"`
	Address        string `json:"address"`
	DestTag        string `json:"dest_tag"`
	TransactionID  string `json:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed"`
	Timeout        int    `json:"timeout"`
	CheckoutURL    string `json:"checkout_url"`
	StatusURL      string `json:"status_url"`
	QrcodeURL      string `json:"qrcode_url"`
}

// CreatePaymentRequest are the parameters for creating a payment
type CreatePaymentRequest struct {
	Amount           string `json:"amount"`
	OriginalCurrency string `json:"currency1"`
	SendingCurrency  string `json:"currency2"`
	BuyerEmail       string `json:"buyer_email"`
}

// PaymentIpnRequest is the response we expect back from the server when the command is "api"
type PaymentIpnRequest struct {
	Status           string `json:"status"`
	StatusText       string `json:"status_text"`
	TxnID            string `json:"txn_id"`
	Currency1        string `json:"currency1"`
	Currency2        string `json:"currency2"`
	Amount1          string `json:"amount1"`
	Amount2          string `json:"amount2"`
	Fee              string `json:"fee"`
	BuyerName        string `json:"buyer_name"`
	Email            string `json:"email"`
	ItemName         string `json:"item_name"`
	ItemNumber       string `json:"item_number"`
	Invoice          string `json:"invoice"`
	Custom           string `json:"custom"`
	SendTX           string `json:"send_tx"` // the tx id of the payment to the merchant. only included when 'status' >= 100 and the payment mode is set to ASAP or nightly or if the payment is paypal passthru
	ReceivedAmount   string `json:"received_amount"`
	ReceivedConfirms string `json:"received_confirms"`
}
