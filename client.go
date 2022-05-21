package coinpayments

import (
	"context"
	"crypto/hmac"
	"crypto/sha512"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	strconv "strconv"
	"strings"
)

type service struct {
	client *Client
}

// Client is the coinpayments API client.
// Do not instantiate this client with Client{}. Use the New method instead.
type Client struct {
	httpClient *http.Client
	common     service
	baseURL    string
	apiKey     string
	apiSecret  string
	version    string

	Payment *paymentService
}

// New creates and returns a new Client from a slice of Option.
func New(options ...Option) *Client {
	config := defaultClientConfig()

	for _, option := range options {
		option.apply(config)
	}

	client := &Client{
		apiKey:     config.apiKey,
		version:    config.version,
		apiSecret:  config.apiSecret,
		httpClient: config.httpClient,
		baseURL:    config.baseURL,
	}

	client.common.client = client
	client.Payment = (*paymentService)(&client.common)
	return client
}

// newRequest creates an API request. A relative URL can be provided in uri,
// in which case it is resolved relative to the BaseURL of the Client.
// URI's should always be specified without a preceding slash.
func (client *Client) newRequest(ctx context.Context, method, cmd string, body url.Values) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, client.baseURL, strings.NewReader(body.Encode()))
	if err != nil {
		return nil, err
	}

	client.addURLParams(req, map[string]string{
		"cmd":     cmd,
		"key":     client.apiKey,
		"format":  "json",
		"version": client.version,
	})

	// generate hmac hash of data and private key
	hash, err := client.computeHMAC(body.Encode())
	if err != nil {
		return nil, err
	}

	req.Header.Add("HMAC", hash)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(body.Encode())))

	return req, nil
}

// addURLParams adds urls parameters to an *http.Request
func (client *Client) addURLParams(request *http.Request, params map[string]string) *http.Request {
	q := request.URL.Query()
	for key, value := range params {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
	return request
}

// do carries out an HTTP request and returns a Response
func (client *Client) do(req *http.Request) (*Response, error) {
	if req == nil {
		return nil, fmt.Errorf("%T cannot be nil", req)
	}

	httpResponse, err := client.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func() { _ = httpResponse.Body.Close() }()

	resp, err := client.newResponse(httpResponse)
	if err != nil {
		return resp, err
	}

	_, err = io.Copy(ioutil.Discard, httpResponse.Body)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

// newResponse converts an *http.Response to *Response
func (client *Client) newResponse(httpResponse *http.Response) (*Response, error) {
	if httpResponse == nil {
		return nil, fmt.Errorf("%T cannot be nil", httpResponse)
	}

	resp := new(Response)
	resp.HTTPResponse = httpResponse

	buf, err := ioutil.ReadAll(resp.HTTPResponse.Body)
	if err != nil {
		return nil, err
	}
	resp.Body = &buf

	return resp, resp.Error()
}

// computeHMAC returns our hmac because on the secret key of our account
func (client *Client) computeHMAC(data string) (string, error) {
	hash := hmac.New(sha512.New, []byte(client.apiSecret))
	if _, err := hash.Write([]byte(data)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}
