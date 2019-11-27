package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"pkg/mod/github.com/moneybin/moneybin-paz/models"
)

func handleGETDonations(w http.ResponseWriter, req *http.Request) {

	donations, err := models.ReadDonations()
	if err != nil {
		fmt.Println("error")
	}

	resp, err := json.Marshal(&donations)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
