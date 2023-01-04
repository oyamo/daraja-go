package darajago

// LipaNaMpesaOnline is a function that initiates a Lipa Na Mpesa Online payment.
// It takes in a LipaNaMpesa struct representing the payment configuration,
// and returns a LipaNaMpesaResponse struct representing the response from the Lipa Na Mpesa API,
// or an ErrorResponse struct representing an error that occurred during the request.
func (d *DarajaApi) LipaNaMpesaOnline(mpesaConfig LipaNaMpesa) (*LipaNaMpesaResponse, *ErrorResponse) {
	// marshal the struct into a map
	payload := struct2Map(mpesaConfig)
	secureResponse, err := performSecurePostRequest[LipaNaMpesaResponse](payload, endpointLipaNaMpesa, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}
