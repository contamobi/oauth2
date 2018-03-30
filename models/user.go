package models

type UserInterface interface {
	GetID() string
	GetClientID() string
	GetAccountID() string
	GetSecret() string
	GetUsername() string
	GetPassword() string
}
