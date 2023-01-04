package darajago

type ErrorResponse struct {
	error
	RequestID    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
}

type B2CConfig struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	SenderIdentifierType   string `json:"SenderIdentifierType"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	Amount                 string `json:"Amount"`
	PartyA                 string `json:"PartyA"`
	PartyB                 string `json:"PartyB"`
	AccountReference       string `json:"AccountReference"`
	Remarks                string `json:"Remarks"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	ResultURL              string `json:"ResultURL"`
}

type B2CResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

type ReversalConfig struct {
	Initiator              string `json:"Initiator"`
	SecurityCredential     string `json:"SecurityCredential"`
	CommandID              string `json:"CommandID"`
	TransactionID          string `json:"TransactionID"`
	Amount                 string `json:"Amount"`
	ReceiverParty          string `json:"ReceiverParty"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	ResultURL              string `json:"ResultURL"`
	QueueTimeOutURL        string `json:"QueueTimeOutURL"`
	Remarks                string `json:"Remarks"`
	Occasion               string `json:"Occasion"`
}

// ReversalResponse represents a response from the Mpesa Reversal.
type ReversalResponse struct {
	// OriginatorConversationID is the unique request ID for tracking a transaction.
	OriginatorConversationID string `json:"OriginatorConversationID"`

	// ConversationID is the unique request ID returned by mpesa for each request made.
	ConversationID string `json:"ConversationID"`

	// ResponseDescription is the response description message.
	ResponseDescription string `json:"ResponseDescription"`
}

type RegisterURLConfig struct {
	ShortCode       string `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}

type TransactionStatus struct {
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	TransactionID      string `json:"TransactionID"`
	PartyA             string `json:"PartyA"`
	IdentifierType     string `json:"IdentifierType"`
	ResultURL          string `json:"ResultURL"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	Remarks            string `json:"Remarks"`
	Occasion           string `json:"Occasion"`
}

// C2BSimulation Simulate paying bills online
type C2BSimulation struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	Msisdn        string `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

// MerchantStatus is used to check the status of a transaction on Lipa Na M-Pesa Online Payment
type MerchantStatus struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}

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

// BalanceQuery is used to query the balance of an M-Pesa account
type BalanceQuery struct {
	Initiator          string `json:"Initiator"`
	SecurityCredential string `json:"SecurityCredential"`
	CommandID          string `json:"CommandID"`
	PartyA             string `json:"PartyA"`
	IdentifierType     string `json:"IdentifierType"`
	Remarks            string `json:"Remarks"`
	QueueTimeOutURL    string `json:"QueueTimeOutURL"`
	ResultURL          string `json:"ResultURL"`
}

// C2BURLRegistration is used to register the confirmation and validation URLs
type C2BURLRegistration struct {
	ShortCode       string `json:"ShortCode"`
	ResponseType    string `json:"ResponseType"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ValidationURL   string `json:"ValidationURL"`
}
