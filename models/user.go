package models

type UserInterface interface {
	GetID() string
	GetClientID() string
	GetAccountID() int64
	GetSecret() string
	GetUsername() string
	GetPassword() string
}
