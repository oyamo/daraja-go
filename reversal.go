package darajago

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

func (d *DarajaApi) ReverseTransaction(transation ReversalConfig) (*ReversalResponse, *ErrorResponse) {
	transation.CommandID = "TransactionReversal"
	// marshal the struct into a map
	payload := struct2Map(transation)
	secureResponse, err := performSecurePostRequest[*ReversalResponse](payload, endpointReversal, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
