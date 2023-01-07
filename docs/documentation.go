package docs

import "github.com/oyamo/daraja-go"

func DOcument() {
	const consumerKey = "E22yMhs"
	const consumerSecret = "zAFGe5cWKv3U1HQ7"

	// Create a new darajago instance
	daraja := darajago.NewDarajaApi(consumerKey, consumerSecret, darajago.ENVIRONMENT_SANDBOX)

	lnmPayload := darajago.LipaNaMpesaPayload{
		BusinessShortCode: "",
		Password:          "",
		Timestamp:         "",
		TransactionType:   "",
		Amount:            "",
		PartyA:            "",
		PartyB:            "",
		PhoneNumber:       "",
		CallBackURL:       "",
		AccountReference:  "",
		TransactionDesc:   "",
	}

	paymentResponse, err := daraja.MakeSTKPushRequest(lnmPayload)
	if err != nil {
		// Handle error
	}

}
