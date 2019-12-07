package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/donations", handleGETDonations)
	http.Handle("/donate", AuthorisedEndpoint(handlePOSTMakeDonation))
	http.Handle("/payment_intent", AuthorisedEndpoint(handleGetPaymentIntent))
	http.Handle("/log_donation", AuthorisedEndpoint(HandleLogDonation))
	http.HandleFunc("/google_login", handleGoogleLogin)
}
