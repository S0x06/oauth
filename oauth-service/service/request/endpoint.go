package request

import (
	"context"
	"log"

	"github.com/go-kit/kit/endpoint"
	"github.com/oauth/v2/service/model"

	"gopkg.in/oauth2.v3/errors"
	"gopkg.in/oauth2.v3/manage"
	"gopkg.in/oauth2.v3/models"
	"gopkg.in/oauth2.v3/server"
	"gopkg.in/oauth2.v3/store"
)

type AuthService interface {
	GetToken(tokenRequest) (tokenResponse, error)
}

type TokenService struct {
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

	return tokenResponse(), nil
}
