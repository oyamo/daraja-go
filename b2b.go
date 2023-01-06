package darajago

// The Business to Business (B2B) API is used to transfer money from one business to another business.
//This API enables the business to pay other businesses.

// B2BPayload represents a request payload for the B2C API.
type B2BPayload struct {
	// InitiatorName is the initiator name.
	InitiatorName string `json:"InitiatorName"`

	// SecurityCredential is the security credential.
	SecurityCredential string `json:"SecurityCredential"`

	// CommandID is the command ID.
	CommandID string `json:"CommandID"`

	// Amount is the amount to be transferred.
	Amount string `json:"Amount"`

	// PartyA is the party A (the organization making the payment).
	PartyA string `json:"PartyA"`

	// PartyB is the party B (the customer receiving the payment).
	PartyB string `json:"PartyB"`

	// Remarks are any remarks for the request.
	Remarks string `json:"Remarks"`

	// QueueTimeOutURL is the queue timeout URL.
	QueueTimeOutURL string `json:"QueueTimeOutURL"`

	// ResultURL is the result URL.
	ResultURL string `json:"ResultURL"`

	// Occasion is the occasion for the payment.
	Occasion string `json:"Occasion"`
}

// B2BResponse  is the response from the C2B API
type B2BResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

// MakeB2BPayment makes a B2C payment.
func (d *DarajaApi) MakeB2BPayment(b2c B2CPayload, certPath string) (*B2CResponse, *ErrorResponse) {
	b2c.CommandID = "BusinessPayment"
	// marshal the struct into a map
	encryptedCredential, err := openSSlEncrypt(b2c.SecurityCredential, certPath)
	if err != nil {
		return nil, &ErrorResponse{error: err, Raw: []byte(err.Error())}
	}
	b2c.SecurityCredential = encryptedCredential

	secureResponse, errRes := performSecurePostRequest[*B2CResponse](b2c, endpointB2CPmtReq, d)
	if err != nil {
		return nil, errRes
	}
	return secureResponse.Body, nil
}
