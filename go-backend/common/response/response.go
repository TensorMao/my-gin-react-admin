package response

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"` // 自定义错误码
	Message string      `json:"msg"`  // 信息
	Data    interface{} `json:"data"` // 数据

}

func response(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	c.JSON(httpCode, Response{
		code,
		msg,
		data,
	})
	return
}

func Success(c *gin.Context, msg string, data interface{}) {
	response(c, http.StatusOK, 20000, msg, data)
}

func Err(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	response(c, httpCode, code, msg, data)
}

const (
	TokenExpired     = "token is expired"
	TokenNotValidYet = "token not active yet"
	TokenMalformed   = "that's not even a token"
	TokenInvalid     = "couldn't handle this token"
)

func TokenFail(c *gin.Context, err error) {
	if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			Err(c, http.StatusUnauthorized, int(jwt.ValidationErrorMalformed), TokenMalformed, nil)
		} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
			// Token is expired
			Err(c, http.StatusUnauthorized, int(jwt.ValidationErrorExpired), TokenExpired, nil)
		} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
			Err(c, http.StatusUnauthorized, int(jwt.ValidationErrorNotValidYet), TokenNotValidYet, nil)
		} else {
			Err(c, http.StatusUnauthorized, 40100, TokenInvalid, nil)
		}
	}
}
