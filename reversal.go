package darajago

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
