package request

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/oauth/v2/service/model"
)

type tokenRequest struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Scope     string `json:"scope"`
}

type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type AuthService interface {
	GetToken(tokenRequest) (tokenResponse, error)
}

type TokenService struct {
}

func (TokenService) GetToken(rq tokenRequest) (tokenResponse, error) {

	var err error

	if rq.AppId == "" {
		return tokenResponse{}, nil
	}

	if rq.AppSecret == "" {
		return tokenResponse{}, nil
	}

	IsExit, err := model.CheckIdentity(rq.AppId, rq.AppSecret)
	if err != nil {
		return tokenResponse{}, nil
	}

	if IsExit != true {
		return tokenResponse{}, nil
	}

	return tokenResponse{
		AccessToken: "DFG",
		ExpiresIn:   "7200",
		Scope:       rq.Scope,
		TokenType:   "Bearer",
	}, nil
}

func MakeGetTokenEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(tokenRequest)
		v, err := svc.GetToken(req)
		if err != nil {
			return v, nil
		}

		return v, nil
	}
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
