package handler

import (
	"encoding/json"
	"net/http"

	"pkg/mod/github.com/moneybin/moneybin-paz/dto"
	"pkg/mod/github.com/moneybin/moneybin-paz/models"
)

func handlePOSTMakeDonation(w http.ResponseWriter, req *http.Request) {

	donation := dto.Donation{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&donation)
	if err != nil {
		panic(err)
	}

	err = models.CreateDonation(&donation)
	if err != nil {
		panic(err)
	}

	resp, err := json.Marshal(&donation)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
