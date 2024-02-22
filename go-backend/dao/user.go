package dao

import (
	"go-backend/global"
	"go-backend/models"
)

func CreateUser(user models.User) error {
	return global.GlobDB.Create(&user).Error
}

// GetUserByID 根据ID获取用户
func GetUserByID(id uint) (*models.User, error) {
	var user models.User
	err := global.GlobDB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByMobile(mobile string) (*models.User, error) {
	var user models.User
	err := global.GlobDB.First(&user, "mobile = ?", mobile).Error

	if err != nil {

		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(user models.User) error {
	return global.GlobDB.Save(&user).Error
}

// DeleteUser 根据ID删除用户
func DeleteUser(id uint) error {
	return global.GlobDB.Delete(&models.User{}, id).Error
}
