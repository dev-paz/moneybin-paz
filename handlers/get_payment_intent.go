package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/stripe/stripe-go/paymentintent"
	stripe "pkg/mod/github.com/stripe/stripe-go"
)

func handleGetPaymentIntent(w http.ResponseWriter, req *http.Request) {

	amount := req.URL.Query().Get("amount")
	if amount == "" {
		return
	}

	a, err := strconv.ParseInt(amount, 10, 64)
	if err != nil {
		return
	}

	stripe.Key = "sk_test_6vNXUZ4qN5uaV4R6LEOUnExS00WkSadUs7"

	params := &stripe.PaymentIntentParams{
		Amount:   &a,
		Currency: stripe.String(string(stripe.CurrencyGBP)),
	}
	pi, err := paymentintent.New(params)
	if err != nil {
		return
	}

	resp, err := json.Marshal(&pi.ClientSecret)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
