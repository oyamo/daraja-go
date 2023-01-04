package darajago

// C2B Simulate paying bills online
type C2B struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	Msisdn        string `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

// C2BResponse is the response from the C2B API
type C2BResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (d *DarajaApi) MakeC2BPayment(c2b C2B) (*C2BResponse, *ErrorResponse) {
	c2b.CommandID = "CustomerPayBillOnline"
	// marshal the struct into a map
	payload := struct2Map(c2b)
	secureResponse, err := performSecurePostRequest[*C2BResponse](payload, endpointSimulatePmtC2B, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}

func (d *DarajaApi) MakeC2BPaymentV2(c2b C2B) (*C2BResponse, *ErrorResponse) {
	c2b.CommandID = "CustomerPayBillOnline"
	// marshal the struct into a map
	payload := struct2Map(c2b)
	secureResponse, err := performSecurePostRequest[*C2BResponse](payload, endpointSimulatePmtC2BV2, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
