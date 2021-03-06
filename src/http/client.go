package http

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// NativeClient provides a concrete implementation, making HTTP requests.
type NativeClient struct {
	Address string
	client  http.Client
}

// NewClient returns a new native Client targeting the provided `address`.
func NewClient(address string) *NativeClient {
	return &NativeClient{
		Address: address,
		client: http.Client{
			Transport: &http.Transport{
				DisableCompression: true,
			},
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// Do returns the Response from executing the provided Request.
func (c *NativeClient) Do(request *Request) (*Response, error) {
	nativeAddress, err := url.Parse(c.Address)
	if err != nil {
		return nil, err
	}

	nativeAddress, err = nativeAddress.Parse(request.Path)
	if err != nil {
		return nil, err
	}

	nativeRequest, err := http.NewRequest(request.Method, nativeAddress.String(), bytes.NewReader(request.Body))
	if err != nil {
		return nil, err
	}

	for key, value := range request.Headers {
		nativeRequest.Header.Set(key, value)
	}
	nativeRequest.Header.Set("User-Agent", "Diplomat/0.0.1")

	nativeResponse, err := c.client.Do(nativeRequest)
	if err != nil {
		return nil, err
	}

	defer nativeResponse.Body.Close()

	response := NewResponse(nativeResponse.StatusCode, strings.Join(strings.Split(nativeResponse.Status, " ")[1:], " "))
	for key, value := range nativeResponse.Header {
		response.Headers[key] = value[0]
	}

	response.Body, err = ioutil.ReadAll(nativeResponse.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
