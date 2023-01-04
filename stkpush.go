package darajago

// LipaNaMpesa is used to initiate a transaction on Lipa Na M-Pesa Online Payment
type LipaNaMpesa struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// LipaNaMpesaResponse represents a response from the Lipa Na Mpesa API.
type LipaNaMpesaResponse struct {
	// MerchantRequestID is the unique request ID for tracking a transaction.
	MerchantRequestID string `json:"MerchantRequestID"`

	// CheckoutRequestID is the unique request ID for the checkout process.
	CheckoutRequestID string `json:"CheckoutRequestID"`

	// ResponseCode is the response code for the request.
	ResponseCode string `json:"ResponseCode"`

	// ResponseDescription is a description of the response.
	ResponseDescription string `json:"ResponseDescription"`

	// CustomerMessage is a message for the customer.
	CustomerMessage string `json:"CustomerMessage"`
}

// LipaNaMpesaOnline is a function that initiates a Lipa Na Mpesa Online payment.
// It takes in a LipaNaMpesa struct representing the payment configuration,
// and returns a LipaNaMpesaResponse struct representing the response from the Lipa Na Mpesa API,
// or an ErrorResponse struct representing an error that occurred during the request.
func (d *DarajaApi) LipaNaMpesaOnline(mpesaConfig LipaNaMpesa) (*LipaNaMpesaResponse, *ErrorResponse) {
	// marshal the struct into a map
	payload := struct2Map(mpesaConfig)
	secureResponse, err := performSecurePostRequest[LipaNaMpesaResponse](payload, endpointLipaNaMpesa, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}
