package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils"
)

func Login(userName string, password string) (fiber.Map, error) {
	// 查找用户是否存在
	var existingUser model.User

	find := database.DB.Where("user_name = ?", userName).First(&existingUser)

	if errors.Is(find.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户不存在")
	}

	if find.Error != nil {
		return nil, errors.New("查询用户失败")
	}

	if !utils.PasswordVerify(password, existingUser.Password) {
		return nil, errors.New("密码错误")
	}

	if existingUser.Status != "1" {
		return nil, errors.New("用户已被禁用")
	}

	token, err := utils.GenerateToken(existingUser.ID, existingUser.UserName, existingUser.UserRoles)

	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return fiber.Map{
		"token":        token,
		"refreshToken": token,
	}, nil
}
