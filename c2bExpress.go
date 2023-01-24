package darajago

// LIPA NA M-PESA ONLINE API also know as M-PESA express
// (STK Push) is a Merchant/Business initiated C2BPayload (Customer to Business) Payment.

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	ExpressDefaultCallBackURL = "daraja-payments/mpesa"
)

// ExpressCallBackFunc A function called after MPESA API sends a request
type ExpressCallBackFunc func(response *CallbackResponse, request *http.Request, err error)

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

// STKPushStatusPayload is used to check the status of a transaction on Lipa Na M-Pesa Online Payment
type STKPushStatusPayload struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}

// STKPushStatusResponse represents a response from the Lipa Na Mpesa API.
type STKPushStatusResponse struct {
	// MerchantRequestID	This is a global unique Identifier for any submited payment request.	String	16813-1590513-1
	MerchantRequestID string `json:"MerchantRequestID"`

	// CheckoutRequestID	This is a global unique Identifier for the processed payment request.	String	ws_CO_DMZ_1234567890
	CheckoutRequestID string `json:"CheckoutRequestID"`

	// ResponseCode This is a Numeric status code that indicates the status of the transaction submission. 0 means successful submission and any other code means an error occured. 	Numeric	0\
	ResponseCode string `json:"ResponseCode"`

	// ResponseDescription	This is a description of the response code.	String	Success. Request accepted for processing
	ResponseDescription string `json:"ResponseDescription"`

	//ResultDesc	Result description is a message from the API that gives the status of the request processing, usualy maps to a specific ResultCode value. It can be a Success processing message or an error description message.	String	E.g: 0: The service request is processed successfully., 1032: Request cancelled by user
	ResultDesc string `json:"ResultDesc"`

	// ResultCode	This is a numeric status code that indicates the status of the transaction processing. 0 means successful processing and any other code means an error occured or the transaction failed.	Numeric	0, 1032
	ResultCode string `json:"ResultCode"`
}

type CallbackResponse struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"Value"`
				} `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
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

// QuerySTKPushStatus is a function that queries the status of a Lipa Na Mpesa Online payment.
// It takes in a STKPushStatusPayload struct representing the payment configuration,
// and returns a STKPushStatusResponse struct representing the response from the Lipa Na Mpesa API,
// or an ErrorResponse struct representing an error that occurred during the request.
func (d *DarajaApi) QuerySTKPushStatus(mpesaConfig STKPushStatusPayload) (*STKPushStatusResponse, *ErrorResponse) {
	//timestamp
	t := time.Now()
	layout := "20060102150405"
	timestamp := t.Format(layout)

	// marshal the struct into a map
	password := base64.StdEncoding.EncodeToString([]byte(mpesaConfig.BusinessShortCode + mpesaConfig.Password + timestamp))

	// add the timestamp and password to the map

	mpesaConfig.Timestamp = timestamp
	mpesaConfig.Password = password

	secureResponse, err := performSecurePostRequest[STKPushStatusResponse](mpesaConfig, endpointQueryLipanaMpesa, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}

// MapExpressGinCallBack Register a callback for listening to MPESA requests
func (d *DarajaApi) MapExpressGinCallBack(gingroup *gin.RouterGroup, callBackUrl string, callback ExpressCallBackFunc) {
	gingroup.POST(callBackUrl, func(context *gin.Context) {
		var callbackResponse CallbackResponse
		err := context.BindJSON(&callbackResponse)
		if err != nil {
			callback(nil, nil, err)
		}
		callback(&callbackResponse, context.Request, nil)
	})
}
