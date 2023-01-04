package darajago

//B2C API is an API used to make payments from a Business to Customers (Pay Outs).
//Also known as Bulk Disbursements. B2C API is used in several scenarios by
//businesses that require to either make Salary Payments, Cashback payments,
//Promotional Payments(e.g. betting winning payouts), winnings, financial
//institutions withdrawal of funds, loan disbursements etc.

// B2CPayload represents a request payload for the B2C API.
type B2CPayload struct {
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

// B2CResponse is the response from the C2B API
type B2CResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

// MakeB2CPayment makes a B2C payment.
func (d *DarajaApi) MakeB2CPayment(b2c B2CPayload) (*B2CResponse, *ErrorResponse) {
	b2c.CommandID = "BusinessPayment"
	// marshal the struct into a map
	payload := struct2Map(b2c)
	secureResponse, err := performSecurePostRequest[*B2CResponse](payload, endpointB2CPmtReq, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
