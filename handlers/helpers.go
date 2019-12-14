package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/moneybin/moneybin-paz/dto"
	oauth2 "google.golang.org/api/oauth2/v1"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

//GenerateJWT creates a new JWT for a user
func GenerateJWT(userID string) (string, error) {
	signingKey := []byte("havealookatbath")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorised"] = true
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).UnixNano()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	}
	return tokenString, nil
}

//GetGoogleUser retrievd a google user object from an id token
func GetGoogleUser(googleToken string) (*dto.GoogleUser, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(googleToken, &dto.GoogleUser{})
	if googleUser, ok := token.Claims.(*dto.GoogleUser); ok {
		return googleUser, nil
	}
	return nil, fmt.Errorf("error getting user %s", err)
}

//GoogleTokenIsValid checks that a google oauth2 token is valid
func GoogleTokenIsValid(token string) error {
	authService, err := oauth2.New(http.DefaultClient)
	if err != nil {
		return err
	}
	// check token is valid
	tokenInfoCall := authService.Tokeninfo()
	tokenInfoCall.IdToken(token)
	ctx, cancelFunc := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancelFunc()
	tokenInfoCall.Context(ctx)
	_, err = tokenInfoCall.Do()
	if err != nil {
		return err
	}
	return nil
}

func SetupCORSResponse(w *http.ResponseWriter, req *http.Request) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Controjwtokenl-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func IsValidJWT(jwtoken string) (bool, *dto.Claims, error) {
	claims := &dto.Claims{}
	// TODO: get this from the right place
	signingKey := []byte("havealookatbath")
	tkn, err := jwt.ParseWithClaims(jwtoken, claims, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})
	if err != nil {
		return false, claims, err
	}
	if !tkn.Valid {
		return false, claims, nil
	}
	return true, claims, nil
}
