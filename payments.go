package coinpayments

type CreatePaymentResponse struct {
	Error  string              `json:"error"`
	Result CreatePaymentResult `json:"result"`
}

type CreatePaymentResult struct {
	Amount         string `json:"amount"`
	Address        string `json:"address"`
	DestTag        string `json:"dest_tag"`
	TxnID          string `json:"txn_id"`
	ConfirmsNeeded string `json:"confirms_needed"`
	Timeout        int    `json:"timeout"`
	CheckoutURL    string `json:"checkout_url"`
	StatusURL      string `json:"status_url"`
	QrcodeURL      string `json:"qrcode_url"`
}

type CreatePaymentRequest struct {
	Amount           string `json:"amount"`
	OriginalCurrency string `json:"currency1"`
	SendingCurrency  string `json:"currency2"`
	BuyerEmail       string `json:"buyer_email"`
}
