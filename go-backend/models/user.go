package models

import (
	"gorm.io/gorm"
	"strconv"
)

const (
	ADMIN  = 0
	NORMAL = 1
)

type User struct {
	gorm.Model
	UserName string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
	Mobile   string `gorm:"unique;not null" binding:"required,mobile"`
	Role     int    `gorm:"not null" binding:"required"`
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID))
}

type PasswordLoginForm struct {
	// 密码  binding:"required"为必填字段,长度大于3小于20
	Password string `form:"password" json:"password" binding:"required,min=3,max=20"`
	//用户名
	Mobile string `form:"mobile" json:"mobile" binding:"required"`

	//Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码
	//CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`       // 验证码id

}
