package util

import (
	//	"fmt"
	//	"encoding/base32"
	"encoding/base64"
	//	"encoding/hex"
	//	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	SecretKey      = "test Jwt"
	TOKEN_TIME_OUT = 2
)

type jwtClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	//	Data interface{}
	AppId     string
	AppSecret string
	//	GrantType string

	TimeOut int64
	//	Key       string
	//	Secret    string
	//	GrantType string
}

func GenerateJwtToken(Issuer string, AppId string, AppSecret string) (base64Token string, expiresIn int, err error) {

	var TimeOut int64 = int64(time.Now().Add(time.Hour * TOKEN_TIME_OUT).Unix())
	//	var TokenTime int64 = int64(time.Now().Add(time.Seconds * TOKEN_TIME_OUT).Unix())

	claims := &jwtClaims{
		jwt.StandardClaims{
			ExpiresAt: TimeOut,
			Issuer:    Issuer,
		},

		AppId,
		AppSecret,

		TimeOut,
	}

	var token string
	JwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = JwtToken.SignedString([]byte(SecretKey))

	expiresIn = TOKEN_TIME_OUT * 60 * 60
	if err != nil {
		return "", expiresIn, err
	}

	//	var hexToken string
	//	hexToken = hex.EncodeToString([]byte(token))
	base64Token = base64.StdEncoding.EncodeToString([]byte(token))

	return base64Token, expiresIn, nil
}

func ParseToken(Base64Token string) ( /*claims jwt.Claims*/ claims interface{}, err error) {

	BytesToken, err := base64.StdEncoding.DecodeString(Base64Token)
	//	BytesToken, err := hex.DecodeString(Base64Token)
	if err != nil {
		return "", err
	}

	Token := string(BytesToken)
	var JwtToken *jwt.Token

	JwtToken, err = jwt.Parse(Token, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	claims = JwtToken.Claims
	return
}
