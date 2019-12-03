package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moneybin/moneybin-paz/dto"
	"github.com/moneybin/moneybin-paz/models"
)

func handleLogin(w http.ResponseWriter, req *http.Request) {
	loginReq := dto.LoginReq{}
	loginResp := dto.LoginResp{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&loginReq)
	if err != nil {
		panic(err)
	}
	user, err := models.ReadUser(loginReq.Email)
	if err != nil {
		fmt.Println("error reading user from email")
		return
	}

	if loginReq.Password != user.Password {
		fmt.Println("incorect password")
		return
	}

	loginResp.Token, err = GenerateJWT(*user)
	if err != nil {
		panic(err)
	}

	loginResp.Authenticated = true

	resp, err := json.Marshal(&loginResp)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
