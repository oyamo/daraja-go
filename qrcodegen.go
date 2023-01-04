package darajago

// QRPayload represents a request payload for a QR code payment.
type QRPayload struct {
	// MerchantName is the merchant name.
	MerchantName string `json:"MerchantName"`

	// RefNo is the reference number.
	RefNo string `json:"RefNo"`

	// Amount is the amount to be paid.
	Amount int `json:"Amount"`

	// TrxCode is the transaction code.
	//Transaction Type. The supported types are:
	//BG: Pay Merchant (Buy Goods).
	//WA: Withdraw Cash at Agent Till.
	//PB: Paybill or Business number.
	//SM: Send Money(Mobile number).
	//SB: Sent to Business. Business number CPI in MSISDN format.
	TrxCode string `json:"TrxCode"`

	// CPI Credit Party Identifier. Can be a Mobile Number, Business Number,
	//Agent Till, Paybill or Business number, Merchant Buy Goods.
	CPI string `json:"CPI"`
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
	secureResponse, err := performSecurePostRequest[QRResponse](struct2Map(payload), endpointQrCode, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}
