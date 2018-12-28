package request

import (
	"context"
	"encoding/json"
	"net/http"
)

type TokenRequest struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	//	Scope     string `json:"scope"`
	GrantType string `json:"grant_type"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	//	Scope       string `json:"scope"`
	//	TokenType   string `json:"token_type"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func DecodeGetTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//	fmt.Println(request.AppId)
	return request, nil
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type ModuleRequest struct {
	Module      string `json:"module"`
	AccessToken string `json:"access_token"`
	//	r      *http.Request
}

type ModuleResponse struct {
	Result bool `json:"result"`
}

func DecodeGetModuleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request ModuleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//	fmt.Println(request.AppId)
	return request, nil
}
