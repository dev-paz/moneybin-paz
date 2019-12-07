package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/moneybin/moneybin-paz/dto"
	"github.com/moneybin/moneybin-paz/models"
)

func handleGoogleLogin(w http.ResponseWriter, req *http.Request) {

	SetupCORSResponse(&w, req)

	if (*req).Method == "OPTIONS" {
		return
	}

	createUserReq := dto.CreateUserReq{}
	LoginResp := dto.LoginResp{}
	user := dto.User{}

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&createUserReq)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	// check the oauth token is valid
	err = GoogleTokenIsValid(createUserReq.Token)
	if err != nil {
		fmt.Println("invalid token")
		return
	}

	// fetch google profile info for user
	googleUser, err := GetGoogleUser(createUserReq.Token)
	if err != nil {
		fmt.Println("Error fetching google profiledytrsj")
		return
	}
	// check if user already exists and sign in if so
	_, err = models.ReadUser(googleUser.Sub)
	if err == nil {
		LoginResp.Token, err = GenerateJWT(user)
		if err != nil {
			fmt.Println("Error generating token")
			return
		}
		LoginResp.Authenticated = true
		resp, err := json.Marshal(&LoginResp)
		if err != nil {
			panic(err)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(resp)

		return
	}
	// if no user is found, make a new one
	user.UserName = googleUser.Name
	user.Email = googleUser.Email
	user.UserID = googleUser.Sub
	user.SignUpTimestamp = 4
	user.LastLoggedIn = 5

	err = models.CreateUser(&user)
	if err != nil {
		panic(err)
	}

	LoginResp.Token, err = GenerateJWT(user)
	if err != nil {
		fmt.Println("Error generating token")
		return
	}

	LoginResp.Authenticated = true
	resp, err := json.Marshal(&LoginResp)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}
