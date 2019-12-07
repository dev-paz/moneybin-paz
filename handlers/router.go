package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/donations", handleGETDonations)
	http.Handle("/donate", Authorised(handlePOSTMakeDonation))
	http.Handle("/payment_intent", Authorised(handleGetPaymentIntent))
	http.Handle("/log_donation", Authorised(HandleLogDonation))
	http.HandleFunc("/google_login", handleGoogleLogin)
}
