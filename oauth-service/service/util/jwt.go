package util

import (
	"fmt"
	//	"encoding/base32"
	"encoding/base64"
	//	"encoding/hex"
	"log"
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
	Data interface{}

	TimeOut int
	//	Key       string
	//	Secret    string
	//	GrantType string
}

func GenerateJwtToken(Issuer string, Data interface{}) (base64Token string, err error) {

	claims := &jwtClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * TOKEN_TIME_OUT).Unix()),
			Issuer:    Issuer,
		},

		Data,

		TOKEN_TIME_OUT,
	}

	var token string
	JwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = JwtToken.SignedString([]byte(SecretKey))

	if err != nil {
		return "", err
	}

	//	var hexToken string
	//	hexToken = hex.EncodeToString([]byte(token))
	base64Token = base64.StdEncoding.EncodeToString([]byte(token))
	fmt.Println("+++")
	fmt.Println(base64Token)
	fmt.Println("+++")
	return base64Token, nil

}

func ParseToken(Base64Token string) (claims jwt.Claims, err error) {

	BytesToken, err := base64.StdEncoding.DecodeString(Base64Token)
	//	BytesToken, err := hex.DecodeString(Base64Token)
	if err != nil {
		log.Fatalln(err)
	}

	Token := string(BytesToken)
	var JwtToken *jwt.Token

	JwtToken, err = jwt.Parse(Token, func(*jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	claims = JwtToken.Claims
	return
}
