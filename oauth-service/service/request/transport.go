package request

import (
	"context"
	"encoding/json"
	"net/http"
)

type tokenRequest struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Scope     string `json:"scope"`
	GrantType string `json:"grant_type"`
	r         *http.Request
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func DecodeGetTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request tokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//	fmt.Println(request.AppId)
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
