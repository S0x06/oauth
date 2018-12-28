package util

import (
	"encoding/json"
	"log"
	"time"
)

type OAuthToken interface {
	//	grantNewToken(client_id string) (interface)
}

type OAuthClient struct {
	ClientId     string
	ClientSecret string
	GrantType    string
}

type OAuthPassWord struct {
	UserId    string
	PassWord  string
	GrantType string
}

type OAuthClientResponse struct {
	AccessToken string
	ExpiresIn   int64
}

const (
	DEFAULT_TOKEN_TYPE             = "bearer"
	TOKEN_VALIDATE_TIME_IN_SECONDS = 60 * 60 * 2
)

//const TOKEN_VALIDATE_TIME_IN_SECONDS = 60 * 60 * 24 * 365

func GetClientToken(client_id string, client_secret string, grant_type string) (OAuthClientResponse, error) {

	if client_id == "" {
		return OAuthClientResponse{}, nil
	}

	if client_secret == "" {
		return OAuthClientResponse{}, nil
	}

	if grant_type == "" {
		return OAuthClientResponse{}, nil
	}

	if grant_type == "client_credentials" {

		//		code, access_token, expires_in := Redis.Get(client_id, REDIS_CODE_PREFIX)

		//		if access_token == "" {
		//		token := new(OAuthClientResponse)
		//		oauth_token := token.grantNewToken(client_id, grant_type)

		//			Redis.Save(client_id, oauth_token, expires_in)

		//		return oauth_token, nil
		//		}

		response := grantNewToken(client_id)

		client := new(OAuthClient)
		client.ClientId = client_id
		client.ClientSecret = client_secret
		client.GrantType = grant_type
		var time_out = time.Second * TOKEN_VALIDATE_TIME_IN_SECONDS

		data, _ := json.Marshal(client)
		RedisSave(response.AccessToken, data, time_out)

		return OAuthClientResponse{AccessToken: response.AccessToken, ExpiresIn: response.ExpiresIn}, nil

	}

	return OAuthClientResponse{}, nil

}

func GetPassWordToken(user_id string, pass_word string, grant_type string) (OAuthClientResponse, error) {

	if user_id == "" {
		return OAuthClientResponse{}, nil
	}

	if pass_word == "" {
		return OAuthClientResponse{}, nil
	}

	if grant_type == "" {
		return OAuthClientResponse{}, nil
	}

	if grant_type == "password" {

		//		code, access_token, expires_in := Redis.Get(client_id, REDIS_CODE_PREFIX)

		//		if access_token == "" {
		//		token := new(OAuthClientResponse)
		//		oauth_token := token.grantNewToken(client_id, grant_type)

		//			Redis.Save(client_id, oauth_token, expires_in)

		//		return oauth_token, nil
		//		}

		response := grantNewToken(user_id)

		client := new(OAuthPassWord)
		client.UserId = user_id
		client.PassWord = pass_word
		client.GrantType = grant_type

		var time_out = time.Second * TOKEN_VALIDATE_TIME_IN_SECONDS
		data, _ := json.Marshal(client)
		RedisSave(response.AccessToken, data, time_out)

		return OAuthClientResponse{AccessToken: response.AccessToken, ExpiresIn: response.ExpiresIn}, nil

	}

	return OAuthClientResponse{}, nil

}

func TokenExpirationInSeconds() int {
	//	log.Println("token expired date:" + token.Expiration)
	expTime, err := time.ParseInLocation("2006-01-02 15:04:05", "2017-05-11 14:06:06", time.Local)
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

func grantNewToken(client_id string) *OAuthClientResponse {

	//	log.Println(EXP_TIME)
	//	t := time.Unix(EXP_TIME, 0)
	//	date := EXP_TIME.Unix()
	//	date := t.Format("2006-01-02 15:04:05")
	//	date := t.Unix()
	//	log.Println(date)
	var response = new(OAuthClientResponse)

	//	token.ClientId = client_id
	//	access_token := RandomCreateBytes(64)
	//	refresh_token := uuid.NewV4().String()
	//	expires_in := TokenExpirationInSeconds()

	response.AccessToken = RandString(64)
	response.ExpiresIn = TOKEN_VALIDATE_TIME_IN_SECONDS
	//	token.RefreshToken = refresh_token
	//	log.Println("grant new token:" + access_token)
	//	data.SaveToken(OAuthToken)
	return response
}
