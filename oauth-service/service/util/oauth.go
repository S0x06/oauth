package util

import (
	"log"
	"time"
)

type OAuthClient struct {
	ClientId     string
	ClientSecret string
	RedirectUri  string
	Scope        string
}

type OAuthClientResponse struct {
	AccessToken string
	ExpiresIn int
}

type OAuthToken struct {
	AccessToken  string
	ExpiresIn   int
	ClientId     string
	TokenType    string
	TokenScope   string
	RefreshToken string
}

const (
	DEFAULT_TOKEN_TYPE             = "bearer"
	REDIS_CODE_PREFIX              = "OAUTH_CODE_"
	REDIS_CODE_TIMEOUT             = 300000
	TOKEN_VALIDATE_TIME_IN_SECONDS = 60
)

//const TOKEN_VALIDATE_TIME_IN_SECONDS = 60 * 60 * 24 * 365

func Token(client_id string, client_secret string, grant_type string, redirect_uri string) (OAuthTokenResponse, error) {

	if client_id == "" {
		return OAuthClientResponse{}, nil
	}

	if client_secret == "" {
		return OAuthClientResponse{}, nil
	}

	if grant_type == "" {
		return OAuthClientResponse{}, nil
	}

	if code == "" {
		return OAuthClientResponse{}, nil
	}

	if grant_type == "client_credentials" {

		code, access_token, expires_in := Redis.Get(client_id, REDIS_CODE_PREFIX)

		if access_token == "" {
			token = new(OAuthToken)
			oauth_token := token.grantNewToken(client_id, grant_type, redirect_uri)
			
			access_token := token.AccessToken
			expires_in := TokenExpirationInSeconds()
			
			Redis.Save(client_id, oauth_token, expires_in)
			
			
			return OAuthClientResponse{AccessToken:,ExpiresIn:TokenExpirationInSeconds()}, nil
		}

	}

	return OAuthClientResponse{}, nil

}

func (token *OAuthToken) TokenExpirationInSeconds() int {
	log.Println("token expired date:" + token.Expiration)
	expTime, err := time.ParseInLocation("2006-01-02 15:04:05", token.Expiration, time.Local)
	if err != nil {
		log.Println("error while parse time")
		log.Println(err)
		return 0
	}
	eu := expTime.Unix()
	nu := time.Now().Unix()

	gap := eu - nu
	if gap <= 0 {
		return 0
	}
	return int(gap)
}

func (token *OAuthToken)IsTokenExpirated() bool {
	gap := TokenExpirationInSeconds()
	if gap == 0 {
		return true
	}
	return false
}

func (token *OAuthToken)grantNewToken(client_id string, token_type string, token_scope string) *OAuthToken {

//	expTime := time.Now().Unix() + TOKEN_VALIDATE_TIME_IN_SECONDS
//	t := time.Unix(expTime, 0)
//	date := t.Format("2006-01-02 15:04:05")
//	var OAuthToken = new(OAuthToken)

	token.ClientId = client_id
	access_token := uuid.NewV4().String()
	refresh_token := uuid.NewV4().String()
	expires_in := token.TokenExpirationInSeconds()
	
	token.AccessToken = access_token
	token.ExpiresIn = expires_in
	token.TokenScope = token_scope
	token.TokenType = token_type
	token.RefreshToken = refresh_token
	log.Println("grant new token:" + token)
	//	data.SaveToken(OAuthToken)
	return OAuthToken
}
