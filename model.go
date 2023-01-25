package darajago

import (
	"encoding/json"
	"fmt"
)

type ErrorResponse struct {
	RequestID    string `json:"requestId"`
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	error        error  `json:"-"`
	Raw          []byte `json:"-"`
}

func (e ErrorResponse) Error() string {
	if e.error != nil {
		return fmt.Sprintf("%v", e.error)
	}
	if e.ErrorMessage == "" && len(e.Raw) != 0 {
		return string(e.Raw)
	}
	bytes, _ := json.Marshal(e)
	return string(bytes)
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
