package coinpayments

import "net/http"

type clientConfig struct {
	httpClient *http.Client
	version    string
	apiKey     string
	apiSecret  string
	ipnSecret  string
	baseURL    string
}

func defaultClientConfig() *clientConfig {
	return &clientConfig{
		httpClient: http.DefaultClient,
		version:    "1",
		baseURL:    "https://www.coinpayments.net/api.php",
	}
}
