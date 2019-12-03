package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/moneybin/moneybin-paz/dto"
	jwt "gopkg.in/dgrijalva/jwt-go.v3"
)

func GenerateJWT(u dto.User) (string, error) {
	signingKey := []byte("havealookatbath")
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorised"] = true
	claims["username"] = "Jordan Matthews"
	claims["exp"] = time.Now().Add(time.Minute * 30).UnixNano()

	tokenString, err := token.SignedString(signingKey)
	if err != nil {
		fmt.Printf(err.Error())
		return "", err
	}
	return tokenString, nil
}

func Authorised(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	signingKey := []byte("havealookatbath")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("error")
				}
				return signingKey, nil
			})
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte(""))
			}
			if token.Valid {
				endpoint(w, r)
			}
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(""))
		}
	})
}
