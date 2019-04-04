package models

type Token struct {
	Jwt     string
	Refresh string
}

type AuthPayload struct {
	Token      *Token
	User       *User
	BalanceBTC *Balance
}
