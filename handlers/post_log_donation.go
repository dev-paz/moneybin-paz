package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/moneybin/moneybin-paz/dto"
	"github.com/moneybin/moneybin-paz/models"
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

// Set your secret key: remember to change this to your live secret key in production
// See your keys here: https://dashboard.stripe.com/account/apikeys

func HandleLogDonation(w http.ResponseWriter, req *http.Request) {
	fmt.Println("made it here!!!!")
	donation := dto.Donation{}

	stripe.Key = "sk_test_6vNXUZ4qN5uaV4R6LEOUnExS00WkSadUs7"

	const MaxBodyBytes = int64(65536)
	req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
	payload, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	// If you are testing your webhook locally with the Stripe CLI you
	// can find the endpoint's secret by running `stripe listen`
	// Otherwise, find your endpoint's secret in your webhook settings in the Developer Dashboard
	endpointSecret := "whsec_21ogFEfqUs7XFgnEziI54AdbU14YIko8"

	// Pass the request body and Stripe-Signature header to ConstructEvent, along
	// with the webhook signing key.
	event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"),
		endpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
		return
	}

	if err := json.Unmarshal(payload, &event); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse webhook body json: %v\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unmarshal the event data into an appropriate struct depending on its Type
	switch event.Type {
	case "payment_intent.succeeded":
		fmt.Println("PaymentIntent was successful!")

		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		donation.Amount = paymentIntent.Amount
		donation.ID = paymentIntent.ID
		donation.DonationCreatedTimestamp = "2019-10-11"
		donation.UserId = "test_user"
		donation.UserName = "Jordan Mattews"

		err = models.CreateDonation(&donation)
		if err != nil {
			panic(err)
		}

		fmt.Println("PaymentIntent was successful!")

	case "payment_method.attached":
		var paymentMethod stripe.PaymentMethod
		err := json.Unmarshal(event.Data.Raw, &paymentMethod)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		fmt.Println("PaymentMethod was attached to a Customer!")

	default:
		fmt.Fprintf(os.Stderr, "Unexpected event type: %s\n", event.Type)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
