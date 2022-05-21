package coinpayments

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/NdoleStudio/coinpayments-go/internal/helpers"
	"github.com/NdoleStudio/coinpayments-go/internal/stubs"
	"github.com/stretchr/testify/assert"
)

func TestPaymentService_CreateTransaction_Request(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	request := new(http.Request)
	key := "test-api-key"
	secret := "test-api-secret"
	server := helpers.MakeRequestCapturingTestServer(http.StatusOK, stubs.CreatePaymentsOkResponse(), request)
	client := New(WithBaseURL(server.URL), WithAPIKey(key), WithAPISecret(secret))
	params := &CreatePaymentRequest{
		Amount:           "1.00000000",
		OriginalCurrency: "USD",
		SendingCurrency:  "USD",
		BuyerEmail:       "john@example.com",
	}

	expectedHMAC := "e3dc88dea6ead6e8e7cef2f1aecec0a1a66f4ada8af0289d03b04e73ed1e218e842ce553938e1c9461767c5459ddac5c611ca3e9622c2511e156d71d097b9563"

	expectedRequest := url.Values{}
	expectedRequest.Add("amount", params.Amount)
	expectedRequest.Add("currency1", params.OriginalCurrency)
	expectedRequest.Add("currency2", params.SendingCurrency)

	// Act
	_, _, err := client.Payment.CreateTransaction(context.Background(), params)

	// Assert
	assert.Nil(t, err)
	assert.NotNil(t, request)

	assert.Equal(t, expectedHMAC, request.Header.Get("HMAC"))

	buf, err := ioutil.ReadAll(request.Body)
	assert.NoError(t, err)

	assert.Equal(t, expectedRequest.Encode(), string(buf))

	// Teardown
	server.Close()
}

func TestPaymentService_CreateTransaction_Ok(t *testing.T) {
	// Setup
	t.Parallel()

	// Arrange
	key := "test-api-key"
	secret := "test-api-secret"
	server := helpers.MakeTestServer(http.StatusOK, stubs.CreatePaymentsOkResponse())
	client := New(WithBaseURL(server.URL), WithAPIKey(key), WithAPISecret(secret))
	params := &CreatePaymentRequest{
		Amount:           "1.00000000",
		OriginalCurrency: "USD",
		SendingCurrency:  "USD",
		BuyerEmail:       "john@example.com",
	}

	// Act
	transactionResponse, apiResponse, err := client.Payment.CreateTransaction(context.Background(), params)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, apiResponse.HTTPResponse.StatusCode)

	assert.Equal(t, &CreatePaymentResponse{
		Error: "ok",
		Result: CreatePaymentResult{
			Amount:         "1.00000000",
			Address:        "ZZZ",
			DestTag:        "YYY",
			TransactionID:  "XXX",
			ConfirmsNeeded: "10",
			Timeout:        9000,
			CheckoutURL:    "https://www.coinpayments.net/index.php?cmd=checkout&id=XXX&key=ZZZ",
			StatusURL:      "https://www.coinpayments.net/index.php?cmd=status&id=XXX&key=ZZZ",
			QrcodeURL:      "https://www.coinpayments.net/qrgen.php?id=XXX&key=ZZZ",
		},
	}, transactionResponse)

	// Teardown
	server.Close()
}
