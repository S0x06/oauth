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
	ExpiresIn   int    `json:"expires_in"`
	//	Scope       string `json:"scope"`
	//	TokenType   string `json:"token_type"`
}

type ErrResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

/**
*客户端授权
 */
func DecodeGetTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request TokenRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//	fmt.Println(request.AppId)
	return request, nil
}

/**
*模块验证
 */

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

type PassWordRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	//	Scope     string `json:"scope"`
	GrantType string `json:"grant_type"`
	Scope     string `json:"scope"`
}

type PassWordResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	GrantType   string `json:"grant_type"`
	Scope       string `json:"scope"`
}

/**
*密码授权
 */
func DecodeGetPassWordTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request PassWordRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//	fmt.Println(request.AppId)
	return request, nil
}

type CodeUserRequest struct {
	AppId        string `json:"app_id"`
	ResponseType string `json:"response_type"`
	RedirectUri  string `json:"redirect_uri"`
	Scope        string `json:"scope"`
	State        string `json:"state"`
}

/**
*授权码模式
 */

type AuthorizationRequest struct {
	AppId     string `json:"app_id"`
	AppSecret string `json:"app_secret"`
	Code      string `json:"code"`
	GrantType string `json:"grant_type"`
}

type AuthorizationResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Uid          int    `json:"uid"`
	Scope        string `json:"scope"`
}

func DecodeGetAuthorizationTokenRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request AuthorizationRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	//	fmt.Println(request.AppId)
	return request, nil
}

/**
刷新token
*/

type RefreshTokenRequest struct {
	AppId        string `json:"app_id"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Uid          int    `json:"uid"`
	Scope        string `json:"scope"`
}

/**
*简化模式
 */

type CodeRequest struct {
	AppId       string `json:"app_id"`
	AppSecret   string `json:"app_secret"`
	GrantType   string `json:"grant_type"`
	RedirectUri string `json:"redirect_uri"`
	Scope       string `json:"scope"`
	State       string `json:"state"`
}

type CodeResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	GrantType   string `json:"grant_type"`
	Scope       string `json:"scope"`
	State       string `json:"state"`
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
