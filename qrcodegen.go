package darajago

type TransactionType string

const (
	TransactionTypeBuyGoods       TransactionType = "BG"
	TransactionTypePayBill        TransactionType = "PB"
	TransactionTypeWithdraw       TransactionType = "WA" // Withdraw Cash at Agent Till.
	TransactionTypeSendMoney      TransactionType = "SM" // Send Money to a Phone Number.
	TransactionTypeSendtoBusiness TransactionType = "SB" // Send Money to a Business.
)

// QRPayload represents a request payload for a QR code payment.
type QRPayload struct {
	// MerchantName is the merchant name.
	MerchantName string `json:"MerchantName"`

	// RefNo is the reference number.
	RefNo string `json:"RefNo"`

	// Amount is the amount to be paid.
	Amount int `json:"Amount"`

	// TransactionType is the transaction code.
	//Transaction Type. The supported types are:
	//
	//BG: Pay Merchant (Buy Goods).
	//
	//WA: Withdraw Cash at Agent Till.
	//
	//PB: Paybill or Business number.
	//
	//SM: Send Money(Mobile number).
	//
	//SB: Sent to Business. Business number CreditPartyIdentifier in MSISDN format.
	TransactionType TransactionType `json:"TrxCode"`

	// CreditPartyIdentifier Credit Party Identifier. Can be a Mobile Number, Business Number,
	//Agent Till, Paybill or Business number, Merchant Buy Goods.
	CreditPartyIdentifier string `json:"CPI"`
}

// QRResponse represents a response from a QR code payment request.
type QRResponse struct {
	// ResponseCode is the response code.
	ResponseCode string `json:"ResponseCode"`

	// RequestID is the request ID.
	RequestID string `json:"RequestID"`

	// ResponseDescription is the response description.
	ResponseDescription string `json:"ResponseDescription"`

	// QRCode is the QR code as a base64-encoded string.
	QRCode string `json:"QRCode"`
}

// MakeQRCodeRequest is a function that generates a QR code for a payment.
// It takes in a QRPayload struct representing the payment configuration,
// and returns a QRResponse struct representing the response from the QR code generation API,
// or an ErrorResponse struct representing an error that occurred during the request.
func (d *DarajaApi) MakeQRCodeRequest(payload QRPayload) (*QRResponse, *ErrorResponse) {
	secureResponse, err := performSecurePostRequest[QRResponse](payload, endpointQrCode, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}
