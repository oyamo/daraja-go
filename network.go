package darajago

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type networkPackage struct {
	Payload io.Reader
	Endpoint string
	Method string
	Headers map[string]string

}

type networkResponse[T any] struct {
	Body T
	StatusCode int
}

func newPackage(payload map[string] interface{}, endpoint string, method string, headers map[string]string) *networkPackage {
	var payloadReader io.Reader
	if method == http.MethodPost {
		payloadBytes, _ := json.Marshal(payload)
		payloadReader = bytes.NewReader(payloadBytes)
	} else if method == http.MethodGet {
		q := url.Values{}
		for key, value := range payload {
			q.Add(key, value.(string))
		}
		payloadReader = bytes.NewReader([]byte(q.Encode()))
	}

	return &networkPackage{
		Payload: payloadReader,
		Endpoint: endpoint,
		Method: method,
		Headers: headers,
	}
}

func newRequest[T any](pac *networkPackage) (*networkResponse[T], error) {
	res := &networkResponse[T]{}
	client := &http.Client{}
	
	req, err := http.NewRequest(pac.Method, pac.Endpoint, nil)
	if err != nil {
		return res, err
	}

	for key, value := range pac.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		return res, err
	}

	defer resp.Body.Close()

	res.StatusCode = resp.StatusCode

	if err := json.NewDecoder(resp.Body).Decode(&res.Body); err != nil {
		return res, err
	}

	return res, nil

}
