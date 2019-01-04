package request

import (
	"context"
	"fmt"
	//	"fmt"
	//	"time"
	//	"log"
	//	"encoding/json"
	//	"fmt"
	//	"crypto/md5"
	//	"encoding/hex"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"

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

/**
*客户端token
 */
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

	if rq.GrantType != "client_credentials" {
		return ErrResponse{Code: 404, Message: "授权模式错误！"}, nil
	}

	IsExit, err := model.CheckIdentity(rq.AppId, rq.AppSecret)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取商户失败！"}, nil
	}

	if IsExit != true {
		return ErrResponse{Code: 404, Message: "用户不存在或者密码错误！"}, nil
	}

	access_token, expires_in, err := util.GenerateJwtToken(rq.GrantType, rq.AppId, rq.AppSecret)

	//	response, err := util.GetClientToken(rq.AppId, rq.AppSecret, rq.GrantType)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取access_token失败！"}, nil
	}

	response := new(TokenResponse)
	response.AccessToken = access_token
	response.ExpiresIn = expires_in
	return response, nil
	//	return TokenResponse{AccessToken: response.AccessToken, ExpiresIn: response.ExpiresIn}, nil
}

/**
*密码token
 */
type PassWordService interface {
	GetPassWordToken(PassWordRequest) (interface{}, error)
	//	GetModule(ModuleRequest) (interface{}, error)
}

type PassWordTokenService struct {
}

func MakeGetPassWordTokenEndpoint(svc PassWordService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(PassWordRequest)
		v, err := svc.GetPassWordToken(req)
		if err != nil {
			return v, nil
		}

		return v, nil
	}
}

func (PassWordTokenService) GetPassWordToken(rq PassWordRequest) (interface{}, error) {

	var err error

	if rq.UserName == "" {
		return ErrResponse{Code: 404, Message: "user_name 不能为空！"}, nil
	}

	if rq.PassWord == "" {
		return ErrResponse{Code: 404, Message: "pass_word 不能为空！"}, nil
	}

	if rq.GrantType != "password" {
		return ErrResponse{Code: 404, Message: "授权模式错误！"}, nil
	}

	//	//	password :=
	//	hasPassword := md5.Sum([]byte(rq.PassWord))
	//	md5Password := fmt.Sprintf("%x", hasPassword) //将[]byte转成16进制

	user, err := model.GetUser(rq.UserName, rq.PassWord)
	if err != nil {
		return ErrResponse{Code: 404, Message: "用户不存在！"}, nil
	}

	if user.ID < 1 {
		return ErrResponse{Code: 404, Message: "用户不存在或者密码错误！"}, nil
	}

	access_token, expires_in, err := util.GenerateJwtToken(rq.GrantType, rq.UserName, rq.PassWord)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取access_token失败！"}, nil
	}

	response := new(PassWordResponse)
	response.AccessToken = access_token
	response.ExpiresIn = expires_in
	response.GrantType = rq.GrantType
	response.Scope = rq.Scope

	return response, nil
}

/**
* 授权码模式
 */
type AuthorizationService interface {
	GetAuthorizationToken(AuthorizationRequest) (interface{}, error)
}

type AuthorizationTokenService struct {
}

func MakeGetAuthorizationTokenEndpoint(svc AuthorizationService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AuthorizationRequest)
		v, err := svc.GetAuthorizationToken(req)
		if err != nil {
			return v, nil
		}

		return v, nil
	}
}

func (AuthorizationTokenService) GetAuthorizationToken(rq AuthorizationRequest) (interface{}, error) {

	var err error

	if rq.AppId == "" {
		return ErrResponse{Code: 404, Message: "app_id 不能为空"}, nil
	}

	if rq.AppSecret == "" {
		return ErrResponse{Code: 404, Message: "app_secret 不能为空"}, nil
	}

	if rq.Code == "" {
		return ErrResponse{Code: 404, Message: "code 不能为空"}, nil
	}

	if rq.GrantType != "authorization_code" {
		return ErrResponse{Code: 404, Message: "authorization_code 错误"}, nil
	}

	appId := "test"
	if rq.AppId != appId {
		return ErrResponse{Code: 404, Message: "app_id 错误"}, nil
	}

	//rq.Code   查询redis
	uid := 123
	scope := "SCOPE"
	refresh_token := uuid.NewV4().String()

	//	if err != nil {
	//		fmt.Printf("Something went wrong: %s", err)
	//		return
	//	}

	fmt.Printf("UUIDv4: %s\n", refresh_token)

	// 获取
	AppSecret := "test"
	access_token, expires_in, err := util.GenerateJwtToken(rq.GrantType, rq.AppId, AppSecret)
	if err != nil {
		return ErrResponse{Code: 404, Message: "获取access_token失败！"}, nil
	}

	response := new(AuthorizationResponse)
	response.AccessToken = access_token
	response.ExpiresIn = expires_in
	response.RefreshToken = refresh_token
	response.Uid = uid
	response.Scope = scope

	return response, nil
}

/**
*
 *   刷新  refresh_token
*/

type RefreshService interface {
	GetRefreshToken(RefreshTokenRequest) (interface{}, error)
}

type RefreshTokenService struct {
}

func MakeGetRefreshTokenEndpoint(svc RefreshService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(RefreshTokenRequest)
		v, err := svc.GetRefreshToken(req)
		if err != nil {
			return v, nil
		}

		return v, nil
	}
}

func (RefreshTokenService) GetRefreshToken(rq RefreshTokenRequest) (interface{}, error) {

	//	var err error

	if rq.AppId == "" {
		return ErrResponse{Code: 404, Message: "app_id 不能为空"}, nil
	}

	if rq.GrantType != "refresh_token" {
		return ErrResponse{Code: 404, Message: "grant_type 错误"}, nil
	}

	if rq.RefreshToken == "" {
		return ErrResponse{Code: 404, Message: "refresh_token 不能为空"}, nil
	}

	response := new(RefreshTokenResponse)

	response.AccessToken = "test"
	response.ExpiresIn = 7200
	response.RefreshToken = "test"
	response.Uid = 1
	response.Scope = "sgfg"

	//	response.GrantType = rq.GrantType
	//	response.Scope = rq.Scope

	return response, nil
}

/**
* 模块认证
 */
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
	//	fmt.Println(rq.AccessToken)
	claims, err := util.ParseToken(rq.AccessToken)

	if err != nil {
		return ErrResponse{Code: 404, Message: "获取 access_token 失败,请重新授权"}, nil
	}

	AppId := claims.(jwt.MapClaims)["AppId"]
	AppSecret := claims.(jwt.MapClaims)["AppSecret"]
	//	TimeOut := claims.(jwt.MapClaims)["TimeOut"]

	if AppId == "" {
		return ErrResponse{Code: 404, Message: "获取 AppId 失败"}, nil
	}
	//	return TimeOut, nil
	if AppSecret == "" {
		return ErrResponse{Code: 404, Message: "获取 AppSecret 失败"}, nil
	}

	//	NowTime := int64(time.Now().Unix())
	//	if NowTime > TimeOut {
	//		return ErrResponse{Code: 404, Message: "获取  access_token 过期, 请重新授权"}, nil
	//	}

	IsExit, err := model.CheckModule(rq.Module)

	if err != nil {
		return ErrResponse{Code: 404, Message: "获取权限失败"}, nil
	}

	if IsExit != true {
		return ModuleResponse{Result: false}, nil
	}

	return ModuleResponse{Result: true}, nil
}
