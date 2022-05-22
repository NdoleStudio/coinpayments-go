package coinpayments

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
)

// paymentService is the API client Payments on coinpayments.net
type paymentService service

// CreateTransaction creates a payment on coinpayments
//
// API Docs: https://www.coinpayments.net/apidoc-create-transaction
func (service *paymentService) CreateTransaction(ctx context.Context, params *CreatePaymentRequest) (*CreatePaymentResponse, *Response, error) {
	payload := url.Values{}
	payload.Add("amount", params.Amount)
	payload.Add("currency1", params.OriginalCurrency)
	payload.Add("currency2", params.SendingCurrency)

	request, err := service.client.newRequest(ctx, http.MethodPost, "create_transaction", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	createPaymentResponse := new(CreatePaymentResponse)
	if err = json.Unmarshal(*response.Body, createPaymentResponse); err != nil {
		return nil, response, err
	}

	return createPaymentResponse, response, nil
}

// GetTransaction fetches information on a transaction
//
// API Docs: https://www.coinpayments.net/apidoc-get-tx-info
func (service *paymentService) GetTransaction(ctx context.Context, transactionID string) (*map[string]interface{}, *Response, error) {
	payload := url.Values{}
	payload.Add("txid", transactionID)

	request, err := service.client.newRequest(ctx, http.MethodPost, "get_tx_info", payload)
	if err != nil {
		return nil, nil, err
	}

	response, err := service.client.do(request)
	if err != nil {
		return nil, response, err
	}

	createPaymentResponse := new(map[string]interface{})
	if err = json.Unmarshal(*response.Body, createPaymentResponse); err != nil {
		return nil, response, err
	}

	return createPaymentResponse, response, nil
}
