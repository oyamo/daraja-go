package docs

import (
	"github.com/oyamo/daraja-go"
)

func DOcument() {
	const consumerKey = "E22yMhs"
	const consumerSecret = "zAFGe5cWKv3U1HQ7"

	// Create a new darajago instance
	daraja := darajago.NewDarajaApi(consumerKey, consumerSecret, darajago.ENVIRONMENT_SANDBOX)

	certPath := "/path/to/cert.pem" // download from safaricom developer portal

	b2bPayload := darajago.B2BPayload{
		InitiatorName:   "oyamosupermarket",
		PassKey:         "200",
		Amount:          "20",
		PartyA:          "",
		PartyB:          "",
		Remarks:         "",
		QueueTimeOutURL: "",
		ResultURL:       "",
		Occasion:        "",
	}

	// initiate b2c payment
	b2cResponse, err := daraja.MakeB2BPayment(b2bPayload, certPath)
	if err != nil {
		// handle error
	}

}
