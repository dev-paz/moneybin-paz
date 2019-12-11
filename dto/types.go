package dto

import jwt "github.com/dgrijalva/jwt-go"

type Donation struct {
	ID                       string
	UserId                   string
	UserName                 string
	Amount                   int64
	DonationCreatedTimestamp string
}

type User struct {
	UserID          string
	Email           string
	UserName        string
	SignUpTimestamp int64
	LastLoggedIn    int64
}

type UserSession struct {
	UserID       string
	RefreshToken string
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

type LoginResp struct {
	Authenticated bool
	AccessToken   string
	RefreshToken  string
}

type LoginReq struct {
	Email    string
	Password string
}

type CreateUserReq struct {
	Token string
}

// TokenInfo struct
type GoogleUser struct {
	Iss string `json:"iss"`
	// userId
	Sub string `json:"sub"`
	Azp string `json:"azp"`
	// clientId
	Aud string `json:"aud"`
	Iat int64  `json:"iat"`
	// expired time
	Exp int64 `json:"exp"`

	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Picture       string `json:"picture"`
	Local         string `json:"locale"`
	jwt.StandardClaims
}
