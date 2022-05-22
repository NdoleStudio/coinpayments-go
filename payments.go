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
	Status           string `form:"status"`
	StatusText       string `form:"status_text"`
	TxnID            string `form:"txn_id"`
	Currency1        string `form:"currency1"`
	Currency2        string `form:"currency2"`
	Amount1          string `form:"amount1"`
	Amount2          string `form:"amount2"`
	Fee              string `form:"fee"`
	IpnType          string `form:"ipn_type"`
	BuyerName        string `form:"buyer_name"`
	Email            string `form:"email"`
	ItemName         string `form:"item_name"`
	ItemNumber       string `form:"item_number"`
	Invoice          string `form:"invoice"`
	Custom           string `form:"custom"`
	ReceivedAmount   string `form:"received_amount"`
	ReceivedConfirms string `form:"received_confirms"`
}

// IsWaiting returns true when the payment is in the waiting state
func (request PaymentIpnRequest) IsWaiting() bool {
	return request.Status == "0" || request.Status == "1" || request.Status == "2" || request.Status == "3"
}

// IsComplete returns true with the payment is completed
func (request PaymentIpnRequest) IsComplete() bool {
	return request.Status == "100"
}

// IsFailed returns ttrue when the payment is failed
func (request PaymentIpnRequest) IsFailed() bool {
	return request.Status == "-2" || request.Status == "-1"
}

// PaymentTransactionResponse is the response gotten when we fetch a transaction
type PaymentTransactionResponse struct {
	Error  string             `json:"error"`
	Result PaymentTransaction `json:"result"`
}

// PaymentTransaction is the transaction details
type PaymentTransaction struct {
	TimeCreated       int    `json:"time_created"`
	TimeExpires       int    `json:"time_expires"`
	Status            int    `json:"status"`
	StatusText        string `json:"status_text"`
	Type              string `json:"type"`
	Coin              string `json:"coin"`
	Amount            int    `json:"amount"`
	AmountFormatted   string `json:"amountf"`
	Received          int    `json:"received"`
	ReceivedFormatted string `json:"receivedf"`
	ReceiveConfirms   int    `json:"recv_confirms"`
	PaymentAddress    string `json:"payment_address"`
	TimeCompleted     int    `json:"time_completed"`
	SenderIP          string `json:"sender_ip"`
}
