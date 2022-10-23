package darajago


const (
	auth_endpoint = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
)

type authPayload struct {
	Authorization string `json:"authorization"` // Base64 encoded string of the consumer key and secret
	GrantType     string `json:"grant_type"`    // client_credentials
}

type authResponse struct {
	AccessToken string `json:"access_token"`   // The access token to be used in subsequent API calls
	ExpiresIn   string `json:"expires_in"`     // The number of seconds before the access token expires
}


type Authorization struct {
	authResponse
}


// NewAuthorization creates a new Authorization object

