package darajago

// LIPA NA M-PESA ONLINE API also know as M-PESA express
// (STK Push) is a Merchant/Business initiated C2B (Customer to Business) Payment.

import (
	"encoding/base64"
	"time"
)

// LipaNaMpesaPayload is used to initiate a transaction on Lipa Na M-Pesa Online Payment
type LipaNaMpesaPayload struct {
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

// MakeSTKPushRequest is a function that initiates a Lipa Na Mpesa Online payment.
// It takes in a LipaNaMpesaPayload struct representing the payment configuration,
// and returns a LipaNaMpesaResponse struct representing the response from the Lipa Na Mpesa API,
// or an ErrorResponse struct representing an error that occurred during the request.
func (d *DarajaApi) MakeSTKPushRequest(mpesaConfig LipaNaMpesaPayload) (*LipaNaMpesaResponse, *ErrorResponse) {
	//timestamp
	t := time.Now()
	layout := "20060102150405"
	timestamp := t.Format(layout)

	// marshal the struct into a map
	password := base64.StdEncoding.EncodeToString([]byte(mpesaConfig.BusinessShortCode + mpesaConfig.Password + timestamp))

	// add the timestamp and password to the map

	mpesaConfig.Timestamp = timestamp
	mpesaConfig.Password = password

	secureResponse, err := performSecurePostRequest[LipaNaMpesaResponse](mpesaConfig, endpointLipaNaMpesa, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}
