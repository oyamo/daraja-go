package darajago

import (
	"encoding/base64"
	"net/http"
)

const (
	auth_endpoint = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
)

type authResponse struct {
	AccessToken string `json:"access_token"` // The access token to be used in subsequent API calls
	ExpiresIn   string `json:"expires_in"`   // The number of seconds before the access token expires
}

type Authorization struct {
	authResponse
}

func NewAuthorization(consumerKey, consumerSecret string) (*Authorization, error) {
	auth := &Authorization{}
	authHeader := map[string]string{
		"Authorisation": base64.StdEncoding.EncodeToString([]byte(consumerKey + ":" + consumerSecret)),
	}

	netPackage := newPackage(nil, auth_endpoint, http.MethodGet, authHeader)
	authResponse, err := newRequest[Authorization](netPackage)
	if err != nil {
		return nil, err
	}

	auth = &authResponse.Body
	return auth, nil
}
