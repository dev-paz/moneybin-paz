package handler

import (
	"net/http"
)

func HandleRequests() {
	http.Handle("/donations", Authorised(handleGETDonations))
	http.Handle("/donate", Authorised(handlePOSTMakeDonation))
	http.Handle("/payment_intent", Authorised(handleGetPaymentIntent))
	http.Handle("/log_donation", Authorised(HandleLogDonation))
	http.HandleFunc("/google_login", handleGoogleLogin)
}
