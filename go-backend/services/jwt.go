package services

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"go-backend/global"
	"time"
)

type jwtService struct{}

var JwtService = new(jwtService)

type JwtUser interface {
	GetUid() string
}

type CustomClaims struct {
	jwt.StandardClaims
}

const (
	TokenType    = "bearer"
	AppGuardName = "app"
)

type TokenOutput struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}

func (service *jwtService) CreateToken(GuardName string, user JwtUser) (tokenData TokenOutput, token *jwt.Token, err error) {
	claims := CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + global.GlobConfig.JWT.JwtTTL,
			Id:        user.GetUid(),
			Issuer:    GuardName, // 用于在中间件中区分不同客户端颁发的 token，避免 token 跨端使用
			NotBefore: time.Now().Unix() - 1000,
		},
	}

	token = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(global.GlobConfig.JWT.Key))
	tokenData = TokenOutput{
		tokenStr, int(global.GlobConfig.JWT.JwtTTL), TokenType}
	return

}

func (service *jwtService) GetUserInfo(GuardName string, id uint) (user JwtUser, err error) {
	switch GuardName {
	case AppGuardName:
		return UserService.GetUserByID(id)
	default:
		err = errors.New("guard " + GuardName + " does not exist")
	}
	return
}
