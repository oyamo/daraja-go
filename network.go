package darajago

import (
	"bytes"
	"encoding/json"
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

func newRequest[T any](pac *networkPackage) (*networkResponse[T], *ErrorResponse) {
	res := &networkResponse[T]{}
	client := &http.Client{}

	req, err := http.NewRequest(pac.Method, pac.Endpoint, nil)
	if err != nil {
		return res, &ErrorResponse{error: err}
	}

	for key, value := range pac.Headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)

	if err != nil {
		return res, &ErrorResponse{error: err}
	}

	defer resp.Body.Close()

	res.StatusCode = resp.StatusCode

	//check 4xx or 5xx error
	if res.StatusCode >= 400 {
		// if the body is not empty, then it is an error response
		if resp.ContentLength > 0 {
			var errorResponse ErrorResponse
			err = json.NewDecoder(resp.Body).Decode(&errorResponse)
			if err != nil {
				return res, &ErrorResponse{error: err}
			}
			return res, &errorResponse
		}
	}
	if err := json.NewDecoder(resp.Body).Decode(&res.Body); err != nil {
		return res, &ErrorResponse{error: err}
	}

	return res, nil
}

func performSecurePostRequest[T any](payload map[string]interface{}, endpoint string, d *DarajaApi) (*networkResponse[T], *ErrorResponse) {
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/json"

	if d.authorization.AccessToken == "" {
		_, err := d.Authorize()
		if err != nil {
			return nil, &ErrorResponse{error: err}
		}
	}
	if time.Now().After(d.nextAuthTime) {
		_, err := d.Authorize()
		if err != nil {
			return nil, &ErrorResponse{error: err}
		}
	}

	// attach the authorization header
	headers["Authorization"] = "Bearer " + d.authorization.AccessToken

	// bundle the request into a package
	netPackage := newRequestPackage(payload, endpoint, http.MethodPost, headers, d.environment)
	return newRequest[T](netPackage)
}
