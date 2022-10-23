package darajago

import (
)

type networkRequest struct {
	Payload interface{} 
	Endpoint string
	Method string
	Headers map[string]string

}

type networkResponse struct {
	Body string
	StatusCode int
}

type network interface {
	MakeRequest(request networkRequest) (networkResponse, error)
}

