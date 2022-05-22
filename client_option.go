package coinpayments

import (
	"net/http"
	"strings"
)

// Option is options for constructing a client
type Option interface {
	apply(config *clientConfig)
}

type clientOptionFunc func(config *clientConfig)

func (fn clientOptionFunc) apply(config *clientConfig) {
	fn(config)
}

// WithHTTPClient sets the underlying HTTP client used for API requests.
// By default, http.DefaultClient is used.
func WithHTTPClient(httpClient *http.Client) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if httpClient != nil {
			config.httpClient = httpClient
		}
	})
}

// WithBaseURL set's the base url for the coinpayments API
func WithBaseURL(baseURL string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		if baseURL != "" {
			config.baseURL = strings.TrimRight(baseURL, "/")
		}
	})
}

// WithAPIKey the coinpayments api key
func WithAPIKey(apiKey string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiKey = apiKey
	})
}

// WithAPISecret the coinpayments api secret
func WithAPISecret(apiSecret string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.apiSecret = apiSecret
	})
}

// WithIPNSecret the coinpayments IPN secret
func WithIPNSecret(ipnSecret string) Option {
	return clientOptionFunc(func(config *clientConfig) {
		config.ipnSecret = ipnSecret
	})
}
