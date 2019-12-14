package handler

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/moneybin/moneybin-paz/models"
)

//HandleRequests routes incoming requests
func HandleRequests() {
	http.HandleFunc("/google_login", handleGoogleLogin)
	http.HandleFunc("/donations", handleGETDonations)
	http.Handle("/donate", authorisedEndpoint(handlePOSTMakeDonation))
	http.Handle("/payment_intent", authorisedEndpoint(handleGetPaymentIntent))
	http.Handle("/log_donation", authorisedEndpoint(HandleLogDonation))
	http.Handle("/authorized", authorisedEndpoint(handleGetAuthStatus))
}

func authorisedEndpoint(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("access_token")
		if err != nil {
			if err == http.ErrNoCookie {
				fmt.Println("no cookie")
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		acessToken := c.Value
		tokenIsValid, claims, err := IsValidJWT(acessToken)
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tokenIsValid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// If the access token has expired, check the refresh token
		if claims.TokenExpiry >= time.Now().UnixNano() {
			c, err := r.Cookie("refresh_token")
			if err != nil {
				if err == http.ErrNoCookie {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			userSession, err := models.ReadUserSession(claims.UserID)
			if err != nil {
				if userSession.RefreshToken != c.Value {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}
			}
		}

		newToken, err := GenerateJWT(claims.UserID)
		if err != nil {
			fmt.Println("Error generating access token")
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "access_token",
			Value:   newToken,
			Expires: time.Now().Add(30 * time.Minute),
		})

		// If authorised, forward the request to the endpoint
		endpoint(w, r)
	})
}
