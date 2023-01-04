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
	authorization  Authorization // Authorization token and expiry
	environment    Environment   // Environment to use
	nextAuthTime   time.Time     // Last time the authorization was updated 	// Expiry time in seconds
	ConsumerKey    string        // Consumer key
	ConsumerSecret string        // Consumer secret
}

type darajaApiImp interface {
	Authorize() (*Authorization, error)
	ReverseTransaction(transation ReversalConfig) (*ReversalResponse, *ErrorResponse)
	LipaNaMpesaOnline(mpesaConfig LipaNaMpesa) (*LipaNaMpesaResponse, *ErrorResponse)
}

// Singleton instance of DarajaApi
var darajaApi *DarajaApi

// NewDarajaApi creates a new DarajaApi instance
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
