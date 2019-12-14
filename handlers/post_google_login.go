package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/moneybin/moneybin-paz/dto"
	"github.com/moneybin/moneybin-paz/models"
)

func handleGoogleLogin(w http.ResponseWriter, req *http.Request) {

	SetupCORSResponse(&w, req)
	if (*req).Method == "OPTIONS" {
		return
	}

	createUserReq := dto.CreateUserReq{}
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
		fmt.Println("Error fetching google profile")
		return
	}
	// check if user already exists and sign in if so
	_, err = models.ReadUser(googleUser.Sub)
	if err == nil {
		user.UserID = googleUser.Sub
		err = loginUser(user, w)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		w.WriteHeader(http.StatusOK)
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

	err = loginUser(user, w)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

// loginUser generates tokens, stores the refresh token and sets the http cookies
func loginUser(user dto.User, w http.ResponseWriter) error {
	acessToken, err := GenerateJWT(user.UserID)
	if err != nil {
		fmt.Println("Error generating access token")
		return err
	}

	refreshToken, err := GenerateJWT(user.UserID)
	if err != nil {
		fmt.Println("Error generating refresh token")
		return err
	}

	us := dto.UserSession{
		RefreshToken: refreshToken,
		UserID:       user.UserID,
	}
	err = models.CreateUserSession(&us)
	if err != nil {
		fmt.Println("error creating user session")
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "access_token",
		Value:   acessToken,
		Expires: time.Now().Add(30 * time.Minute),
	})

	http.SetCookie(w, &http.Cookie{
		Name:    "refresh_token",
		Value:   refreshToken,
		Expires: time.Now().Add(72 * time.Hour),
	})

	return nil
}
