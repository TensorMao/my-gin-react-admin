package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go-backend/common/response"
	"go-backend/global"
	"go-backend/services"
	"net/http"
	"strconv"
	"time"
)

func JWTAuth(GuardName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 Authorization 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" {
			response.Err(c, http.StatusUnauthorized, 40100, "Please Login", nil)
			c.Abort()
			return
		}

		tokenStr = tokenStr[len(services.TokenType)+1:]
		// parseToken 解析token包含的信息
		token, err := jwt.ParseWithClaims(tokenStr, &services.CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
			return []byte(global.GlobConfig.JWT.Key), nil
		})
		if err != nil || services.JwtService.IsInBlacklist(tokenStr) {
			//TODO
			response.TokenFail(c, err)
			c.Abort()
			return

		}

		claims := token.Claims.(*services.CustomClaims)

		if claims.Issuer != GuardName {
			//TODO
			response.TokenFail(c, err)
			c.Abort()
			return

		}
		if claims.ExpiresAt-time.Now().Unix() < global.GlobConfig.JWT.RefreshGracePeriod {
			lock := global.Lock("refresh_token_lock"+claims.Id, global.GlobConfig.JWT.JwtBlacklistGracePeriod)
			num, _ := strconv.ParseUint(claims.Id, 10, 32)
			if lock.Get() {
				user, err := services.JwtService.GetUserInfo(GuardName, uint(num))
				if err != nil {
					global.GlobLogger.Error(err.Error())
					lock.Release()
				} else {
					tokenData, _, _ := services.JwtService.CreateToken(GuardName, user)
					c.Header("new-token", tokenData.AccessToken)
					c.Header("new-expires-in", strconv.Itoa(tokenData.ExpiresIn))
					_ = services.JwtService.JoinBlackList(token)
				}
			}
		}

		c.Set("token", token)
		c.Set("id", claims.Id)
	}
}

// ParseToken 解析 token
/*
	token, err :=

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*services.CustomClaims);claims.Issuer!=GuardName{}

		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid

	}

}

// RefreshToken 更新token
func  RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = time.Now().Add(1 * time.Hour).Unix()
		return j.CreateToken(*claims)
	}
	return "", TokenInvalid
}*/
