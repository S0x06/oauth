package request

import (
	"context"
	"fmt"
	//	"log"
	"encoding/json"
	//	"fmt"
	"github.com/dgrijalva/jwt-go"

	"github.com/go-kit/kit/endpoint"
	"github.com/oauth/v2/service/model"
	"github.com/oauth/v2/service/util"
)

type AuthService interface {
	GetToken(TokenRequest) (interface{}, error)
	//	GetModule(ModuleRequest) (interface{}, error)
}

type TokenService struct {
}

func MakeGetTokenEndpoint(svc AuthService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(TokenRequest)
		v, err := svc.GetToken(req)
		if err != nil {
			return v, nil
		}

		return v, nil
	}
}

func (TokenService) GetToken(rq TokenRequest) (interface{}, error) {

	var err error

	if rq.AppId == "" {
		return ErrResponse{Code: 404, Message: "app_id 不能为空！"}, nil
	}

	if rq.AppSecret == "" {
		return ErrResponse{Code: 404, Message: "app_secret 不能为空！"}, nil
	}

	IsExit, err := model.CheckIdentity(rq.AppId, rq.AppSecret)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取商户失败！"}, nil
	}

	if IsExit != true {
		return ErrResponse{Code: 404, Message: "用户不存在或者密码错误！"}, nil
	}

	response, err := util.GenerateJwtToken(rq.GrantType, rq)

	//	response, err := util.GetClientToken(rq.AppId, rq.AppSecret, rq.GrantType)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取access_token失败！"}, nil
	}
	fmt.Println(response)
	fmt.Println("=====")
	return response, nil
	//	return TokenResponse{AccessToken: response.AccessToken, ExpiresIn: response.ExpiresIn}, nil
}

type ModuleService interface {
	GetModule(ModuleRequest) (interface{}, error)
}

type OauthModuleService struct {
}

func MakeGetModuleEndpoint(svc ModuleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(ModuleRequest)
		v, err := svc.GetModule(req)
		if err != nil {
			return v, nil
		}

		return v, nil
	}
}

func (OauthModuleService) GetModule(rq ModuleRequest) (interface{}, error) {

	if rq.AccessToken == "" {
		return ErrResponse{Code: 404, Message: "access_token 请填写"}, nil

	}
	if rq.Module == "" {
		return ErrResponse{Code: 404, Message: "无权限"}, nil
	}
	fmt.Println(rq.AccessToken)
	claims, err := util.ParseToken(rq.AccessToken)

	fmt.Println("claims TimeOut:", claims.(jwt.MapClaims)["TimeOut"])

	//	request := claims.(jwt.MapClaims)["Data"].(TokenRequest)

	fmt.Println("claims TimeOut:", request)

	access_token, data, err := util.RedisGet(rq.AccessToken)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取 access_token 失败"}, nil
	}

	if access_token == "" {
		return ErrResponse{Code: 403, Message: "请重新授权"}, nil
	}

	var client = new(util.OAuthClient)
	err = json.Unmarshal([]byte(data), client)
	if err != nil {
		return ErrResponse{Code: 404, Message: "access_token 请填写"}, nil
	}

	IsExit, err := model.CheckModule(rq.Module)

	if err != nil {
		return ErrResponse{Code: 404, Message: "获取权限失败"}, nil
	}

	if IsExit != true {
		return ModuleResponse{Result: false}, nil
	}

	return ModuleResponse{Result: true}, nil
}
