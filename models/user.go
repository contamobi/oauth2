package models

type UserInterface interface {
	GetID() int64
	GetClientID() string
	GetAccountID() int64
	GetSecret() string
	GetUsername() string
	GetPassword() string
}
