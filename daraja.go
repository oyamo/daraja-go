package darajago

import (
	"time"
)

const (
	// ENVIRONMENT_SANDBOX is the sandbox environment
	ENVIRONMENT_SANDBOX = "sandbox"
	// ENVIRONMENT_PRODUCTION is the production environment
	ENVIRONMENT_PRODUCTION = "production"
)

type Environment string

type DarajaApi struct {
	authorization  Authorization // Authorization token and expiry kept in memory for subsequent requests
	environment    Environment   // Environment to use eg sandbox or production
	nextAuthTime   time.Time     // When to request access token next
	ConsumerKey    string        // Consumer key
	ConsumerSecret string        // Consumer secret
}

type darajaAAPIImpl interface {
	Authorize() (*Authorization, error)
	ReverseTransaction(transation ReversalPayload) (*ReversalResponse, *ErrorResponse)
	MakeSTKPushRequest(mpesaConfig LipaNaMpesaPayload) (*LipaNaMpesaResponse, *ErrorResponse)
	MakeB2BPayment(b2c B2CPayload) (*B2CResponse, *ErrorResponse)
	MakeB2CPayment(b2c B2CPayload) (*B2CResponse, *ErrorResponse)
	MakeQRCodeRequest(payload QRPayload) (*QRResponse, *ErrorResponse)
	MakeC2BPayment(c2b C2BPayload) (*C2BResponse, *ErrorResponse)
	MakeC2BPaymentV2(c2b C2BPayload) (*C2BResponse, *ErrorResponse)
}

// Singleton instance of DarajaApi
var darajaApi *DarajaApi

// NewDarajaApi is a function that creates a new DarajaApi struct.
// It takes a consumer key, a consumer secret, and an Environment as input,
// and returns a pointer to the DarajaApi struct.
func NewDarajaApi(consumerKey, consumerSecret string, environment Environment) *DarajaApi {
	if darajaApi == nil {
		darajaApi = &DarajaApi{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			environment:    environment,
		}
	}
	return darajaApi
}

// Authorize is a function that retrieves an authorization token from the Mpesa API.
// It returns an Authorization struct representing the authorization token,
// or an error if an error occurred during the request.
func (d *DarajaApi) Authorize() (*Authorization, error) {
	authTimeStart := time.Now()
	auth, err := newAuthorization(d.ConsumerKey, d.ConsumerSecret, d.environment)
	if err != nil {
		return nil, err
	}
	expiry, err := time.ParseDuration(auth.ExpiresIn + "s")
	if err != nil {
		return nil, err
	}
	d.authorization = *auth
	d.nextAuthTime = authTimeStart.Add(expiry)

	return auth, nil
}
