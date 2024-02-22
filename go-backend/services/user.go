package services

import (
	"errors"
	"go-backend/common/request"
	"go-backend/dao"
	"go-backend/models"
	"go-backend/utils"
)

type userService struct{}

var UserService = new(userService)

// CreateUser 创建新用户
func CreateUser(user models.User) error {
	return dao.CreateUser(user)
}

// GetUserByID 根据ID获取用户
func (service *userService) GetUserByID(id uint) (*models.User, error) {
	return dao.GetUserByID(id)
}

// UpdateUser 更新用户信息
func UpdateUser(user models.User) error {
	return dao.UpdateUser(user)
}

// DeleteUser 根据ID删除用户
func DeleteUser(id uint) error {
	return dao.DeleteUser(id)
}

func (service *userService) Login(params request.Login) (user *models.User, err error) {
	user, err = dao.GetUserByMobile(params.Mobile)
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("---the user doesn't exist ,or your password is incorrect")
	}
	return
}

func (service *userService) Register(params request.Register) (user models.User, err error) {
	result, err := dao.GetUserByMobile(params.Mobile)
	if result != nil {
		err = errors.New("the mobile has existed")
		return
	}
	user = models.User{UserName: params.Name, Mobile: params.Mobile, Password: utils.BcryptMake([]byte(params.Password)), Role: models.NORMAL}
	err = dao.CreateUser(user)
	return

}
