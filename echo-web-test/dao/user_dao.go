package dao

import (
	"github.com/dzjyyds666/echo-web-test/config"
	"github.com/dzjyyds666/echo-web-test/models"
)

// 创建用户
func CreateUser(user *models.User) error {
	return config.DB.Create(user).Error
}

// GetUserById 根据用户Id获取用户信息
func GetUserById(id string) (*models.User, error) {
	var user models.User
	if err := config.DB.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUser 更新用户信息
func UpdateUser(user *models.User) error {
	return config.DB.Save(user).Error
}

// DeleteUserByID 根据 ID 删除用户
func DeleteUserByID(id uint) error {
	return config.DB.Delete(&models.User{}, id).Error
}
