package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go-backend/common/response"
	"go.uber.org/zap"
	"net/http"
)

var store = base64Captcha.DefaultMemStore

// GetCaptcha 获取验证码
func GetCaptcha(ctx *gin.Context) {
	//
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	// b64s是图片的base64编码
	id, b64s, _, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("generating captcha failed ,:%s ", err.Error())
		response.Err(ctx, http.StatusInternalServerError, 50000, "Failed to generate captcha", nil)
		return
	}
	response.Success(ctx, "generating captcha succeeded", gin.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
