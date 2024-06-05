package controller

import (
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/services"
	"soybean-admin-go/app/utils/response"
)

type UserReq struct {
	// 用户名
	// required: true
	// example: admin
	UserName string `json:"userName"`
	// 密码
	// required: true
	// example: 123456
	Password string `json:"password"`
}

type GetUserResponse struct {
	Buttons  []string `json:"buttons"`
	Roles    []string `json:"roles"`
	UserId   uint     `json:"userId"`
	UserName string   `json:"userName"`
}

func GetUserInfo(ctx *fiber.Ctx) error {
	userId := ctx.Locals("userId").(float64)
	var result GetUserResponse

	findUser, err := services.GetUserInfo(userId)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	result.UserId = findUser.ID
	result.UserName = findUser.UserName
	result.Roles = findUser.UserRoles
	result.Buttons = []string{}

	return response.OkWithData(result, ctx)
}

func GetUserList(ctx *fiber.Ctx) error {
	//  获取请求参数
	current := ctx.QueryInt("current", 1)
	size := ctx.QueryInt("size", 10)
	userName := ctx.Query("userName", "")
	userGender := ctx.Query("userGender", "")
	UserPhone := ctx.Query("userPhone", "")
	UserEmail := ctx.Query("userEmail", "")
	Status := ctx.Query("status", "")
	// 查询用户列表
	result, err := services.GetUserList(userName, userGender, UserPhone, UserEmail, Status, current, size)
	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}
	return response.OkWithData(result, ctx)
}

// CreateUser 创建用户
// @Summary 创建用户
// @Description 创建一个新的系统用户
// @Tags 用户管理
// @Accept json
// @Produce json
// @Param user body UserReq true "用户信息"
// @Success 200 {object} model.User "成功创建用户并返回新创建的用户信息"
// @Router /user [post]
func CreateUser(ctx *fiber.Ctx) error {
	// 获取请求参数
	var userReq UserReq
	// 解析参数
	if err := ctx.BodyParser(&userReq); err != nil {
		return response.FailWithMessage("参数解析错误", ctx)
	}

	// 创建用户
	result, err := services.CreateUser(userReq.UserName, userReq.Password)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	return response.OkWithData(result, ctx)
}
