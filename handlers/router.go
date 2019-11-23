package handler

import (
	"net/http"
)

func HandleRequests() {
	http.HandleFunc("/donations", handleGETDonations)
	http.HandleFunc("/donate", handlePOSTMakeDonation)
}
