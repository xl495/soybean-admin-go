package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils"
	"soybean-admin-go/app/utils/response"
)

func CreateUser(username string, password string) (fiber.Map, error) {
	pws, err := utils.PasswordHash(password)

	if err != nil {
		return nil, errors.New("密码加密错误")
	}

	// 查找用户名是否存在
	var existingUser model.User
	database.DB.Where("username = ?", username).First(&existingUser)

	if existingUser.ID != 0 {
		return nil, errors.New("用户名已存在")
	}

	newUser := &model.User{
		Username: username,
		Password: pws,
	}

	// 获取创建并返回用户数据 使用 bcrypt 加密密码
	dbUser := database.DB.Create(newUser)

	responseData := make(fiber.Map)

	responseData["ID"] = newUser.ID
	responseData["username"] = newUser.Username

	if dbUser.Error != nil {
		return nil, errors.New("创建用户失败")
	}

	return responseData, nil
}

func GetUserList(username string, current int, size int) (response.PageResult, error) {
	var users []model.User
	var total int64

	if username != "" {
		database.DB.Where("username like ?", "%"+username+"%").Offset((current - 1) * size).Limit(size).Find(&users)
	} else {
		database.DB.Offset((current - 1) * size).Limit(size).Find(&users)
	}

	return response.PageResult{
		Records: users,
		Total:   total,
		Current: current,
		Size:    size,
	}, nil

}
