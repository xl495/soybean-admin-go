package services

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils"
	"soybean-admin-go/app/utils/response"
)

func GetUserInfo(userId float64) (model.User, error) {
	var user model.User
	findUser := database.DB.Where("id = ?", userId).First(&user)

	if findUser.Error != nil {
		return model.User{}, errors.New("用户不存在")
	}

	return user, nil
}

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
		UserName: username,
		Password: pws,
	}

	// 获取创建并返回用户数据 使用 bcrypt 加密密码
	dbUser := database.DB.Create(newUser)

	responseData := make(fiber.Map)

	responseData["ID"] = newUser.ID
	responseData["username"] = newUser.UserName

	if dbUser.Error != nil {
		return nil, errors.New("创建用户失败")
	}

	return responseData, nil
}

func AddUser(userName, password, userPhone, userEmail string, userRoles []string) (model.User, error) {
	var user model.User
	user.UserName = userName
	user.UserPhone = userPhone
	user.UserEmail = userEmail
	user.UserRoles = userRoles

	pws, err := utils.PasswordHash(password)

	if err != nil {
		return user, errors.New("密码加密错误")
	}

	user.Password = pws
	user.Status = "1"
	if database.DB.Create(&user).Error != nil {
		return user, errors.New("创建用户失败")
	}
	return user, nil
}

func EditUser(id int, userName, password, userPhone, userEmail string, userRoles []string) (model.User, error) {
	var user model.User
	user.UserName = userName
	user.UserPhone = userPhone
	user.UserEmail = userEmail
	user.UserRoles = userRoles

	pws, err := utils.PasswordHash(password)

	if err != nil {
		return user, errors.New("密码加密错误")
	}

	user.Password = pws

	if database.DB.Model(&user).Where("id = ?", id).Updates(&user).Error != nil {
		return user, errors.New("更新用户失败")
	}
	return user, nil
}

func RemoveUsers(ids []int) error {
	query := database.DB.Delete(&model.User{}, ids)
	if query.Error != nil {
		return query.Error
	}
	return nil
}

func GetUserList(userName, userGender, UserPhone, UserEmail, Status string, current int, size int) (response.PageResult, error) {
	var users []model.User
	var total int64

	// 开始构造查询
	query := database.DB

	// 如果 userName 不为空，则添加用户名过滤条件
	if userName != "" {
		query = query.Where("user_name LIKE ?", "%"+userName+"%")
	}
	// 如果 Status 不为空，则添加用户名过滤条件
	if Status != "" {
		query = query.Where("status = ?", Status)
	}
	// 如果 userGender 不为空，则添加用户名过滤条件
	if userGender != "" {
		query = query.Where("user_gender = ?", userGender)
	}
	// 如果 UserPhone 不为空，则添加用户名过滤条件
	if UserPhone != "" {
		query = query.Where("user_phone LIKE ?", "%"+UserPhone+"%")
	}
	// 如果 UserEmail 不为空，则添加用户名过滤条件
	if UserEmail != "" {
		query = query.Where("user_email LIKE ?", "%"+UserEmail+"%")
	}

	// 使用分页和获取用户列表
	query.Offset((current - 1) * size).Limit(size).Find(&users)

	// 计算总数
	query.Model(&model.User{}).Count(&total)

	// 确保返回的 users 不是 nil
	if users == nil {
		users = []model.User{}
	}

	return response.PageResult{
		Records: users,
		Total:   total,
		Current: current,
		Size:    size,
	}, nil
}
