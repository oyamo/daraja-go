package darajago

import (
	"time"
)

const (
	ENVIRONMENT_SANDBOX    = "sandbox"
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
}

// NewDarajaApi creates a new DarajaApi instance
func NewDarajaApi(consumerKey, consumerSecret string, environment Environment) *DarajaApi {
	return &DarajaApi{
		ConsumerKey:    consumerKey,
		ConsumerSecret: consumerSecret,
		environment:    environment,
	}
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
