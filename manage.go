package oauth2

import (
	"time"
)

// TokenGenerateRequest provide to generate the token request parameters
type TokenGenerateRequest struct {
	ClientID       string
	ClientSecret   string
	UserID         string
	AccountID      string
	RedirectURI    string
	Scope          string
	Code           string
	Refresh        string
	AccessTokenExp time.Duration
}

// Manager authorization management interface
type Manager interface {
	// check the interface implementation
	CheckInterface() (err error)

	// get the client information
	GetClient(clientID string) (cli ClientInfo, err error)

	// generate the authorization token(code)
	GenerateAuthToken(rt ResponseType, tgr *TokenGenerateRequest) (authToken TokenInfo, err error)

	// generate the access token
	GenerateAccessToken(rt GrantType, tgr *TokenGenerateRequest) (accessToken TokenInfo, err error)

	// refreshing an access token
	RefreshAccessToken(tgr *TokenGenerateRequest) (accessToken TokenInfo, err error)

	// use the access token to delete the token information
	RemoveAccessToken(access string) (err error)

	// use the refresh token to delete the token information
	RemoveRefreshToken(refresh string) (err error)

	// according to the access token for corresponding token information
	LoadAccessToken(access string) (ti TokenInfo, err error)

	// according to the refresh token for corresponding token information
	LoadRefreshToken(refresh string) (ti TokenInfo, err error)
}
