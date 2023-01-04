package darajago

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

type networkPackage struct {
	Payload  io.Reader
	Endpoint string
	Method   string
	Headers  map[string]string
}

type networkResponse[T any] struct {
	Body       T
	StatusCode int
}

func newRequestPackage(payload map[string]interface{}, endpoint string, method string, headers map[string]string, env Environment) *networkPackage {
	var payloadReader io.Reader
	var reqUrl = baseUrlSandbox
	if env == ENVIRONMENT_PRODUCTION {
		reqUrl = baseUrlLive
	}
	reqUrl = reqUrl + endpoint
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
		Payload:  payloadReader,
		Endpoint: endpoint,
		Method:   method,
		Headers:  headers,
	}
}

func (p *networkPackage) addHeader(key string, value string) {
	if p.Headers == nil {
		p.Headers = make(map[string]string)
	}
	p.Headers[key] = value
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

	//check 4xx or 5xx error
	if res.StatusCode >= 400 {
		return nil, errors.New(resp.Status)
	}
	if err := json.NewDecoder(resp.Body).Decode(&res.Body); err != nil {
		return res, err
	}

	return res, nil

}

func performSecurePostRequest[T any](payload map[string]interface{}, endpoint string, d *DarajaApi) (*networkResponse[T], error) {
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"

	if d.authorization.AccessToken == "" {
		_, err := d.Authorize()
		if err != nil {
			return nil, err
		}
	}
	if time.Now().After(d.nextAuthTime) {
		_, err := d.Authorize()
		if err != nil {
			return nil, err
		}
	}

	// attach the authorization header
	headers["Authorization"] = "Bearer " + d.authorization.AccessToken

	// bundle the request into a package
	netPackage := newRequestPackage(payload, endpoint, http.MethodPost, headers, d.environment)
	return newRequest[T](netPackage)
}
