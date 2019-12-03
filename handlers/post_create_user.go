package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	guuid "github.com/google/uuid"

	"github.com/moneybin/moneybin-paz/dto"
	"github.com/moneybin/moneybin-paz/models"
)

func handleCreateUser(w http.ResponseWriter, req *http.Request) {

	fmt.Println("hello")
	user := dto.User{}
	createUserResp := dto.CreateUserRsp{}

	fmt.Println(req.Body)

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	fmt.Println(user)

	user.UserName = "Jordan Matthews"
	user.UserID = guuid.New().String()
	user.SignUpTimestamp = 4
	user.LastLoggedIn = 5

	err = models.CreateUser(&user)
	if err != nil {
		panic(err)
	}

	createUserResp.Token, err = GenerateJWT(user)
	if err != nil {
		fmt.Println("Error generating token")
		return
	}

	createUserResp.User = user
	resp, err := json.Marshal(&createUserResp)
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
