package dto

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
	Password        string
	UserName        string
	SignUpTimestamp int64
	LastLoggedIn    int64
}

type CreateUserRsp struct {
	User  User
	Token string
}

type LoginResp struct {
	Authenticated bool
	Token         string
}

type LoginReq struct {
	Email    string
	Password string
}
