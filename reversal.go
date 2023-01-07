package darajago

type ReversalPayload struct {
	Initiator              string `json:"Initiator"`
	PassKey                string `json:"SecurityCredential"`
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

func (d *DarajaApi) ReverseTransaction(transation ReversalPayload, certPath string) (*ReversalResponse, *ErrorResponse) {
	transation.CommandID = "TransactionReversal"
	// marshal the struct into a map
	encryptedCredential, err := openSSlEncrypt(transation.PassKey, certPath)
	if err != nil {
		return nil, &ErrorResponse{error: err, Raw: []byte(err.Error())}
	}
	transation.PassKey = encryptedCredential

	secureResponse, errRes := performSecurePostRequest[*ReversalResponse](transation, endpointB2CPmtReq, d)
	if errRes != nil {
		return nil, errRes
	}
	return secureResponse.Body, nil
}
