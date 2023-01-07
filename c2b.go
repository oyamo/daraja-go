package darajago

// C2BPayload Simulate paying bills online
type C2BPayload struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	Msisdn        string `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

// C2BResponse is the response from the C2BPayload API
type C2BResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

// C2BRegistrationPayload is the payload for C2BPayload shortcode registration.
type C2BRegistrationPayload struct {
	// ValidationURL is the URL that receives the validation request from the API upon payment submission.
	// The validation URL is only called if the external validation on the registered shortcode is enabled.
	ValidationURL string `json:"ValidationURL"`

	// ConfirmationURL is the URL that receives the confirmation request from the API upon payment completion.
	ConfirmationURL string `json:"ConfirmationURL"`

	// ResponseType specifies what is to happen if the validation URL is not reachable.
	// Only two values are allowed: Completed or Cancelled.
	// Completed means MPesa will automatically complete the transaction,
	// whereas Cancelled means MPesa will automatically cancel the transaction.
	ResponseType string `json:"ResponseType"`

	// ShortCode is the short code of the organization.
	ShortCode string `json:"ShortCode"`
}

// C2BRegistrationResponse is the response from the C2BPayload API
type C2BRegistrationResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	// ResponseDescription 	This is the status of the request.
	ResponseDescription string `json:"ResponseDescription"`
}

func (d *DarajaApi) RegisterC2BCallback(payload C2BRegistrationPayload) (*C2BRegistrationResponse, *ErrorResponse) {
	secureResponse, err := performSecurePostRequest[*C2BRegistrationResponse](payload, endpointRegisterConfirmValidation, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}

func (d *DarajaApi) MakeC2BPayment(c2b C2BPayload) (*C2BResponse, *ErrorResponse) {
	c2b.CommandID = "CustomerPayBillOnline"
	// marshal the struct into a map
	secureResponse, err := performSecurePostRequest[*C2BResponse](c2b, endpointSimulatePmtC2B, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}

func (d *DarajaApi) MakeC2BPaymentV2(c2b C2BPayload) (*C2BResponse, *ErrorResponse) {
	c2b.CommandID = "CustomerPayBillOnline"
	// marshal the struct into a map
	secureResponse, err := performSecurePostRequest[*C2BResponse](c2b, endpointSimulatePmtC2BV2, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
