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
	Username string `json:"username"`
	// 密码
	// required: true
	// example: 123456
	Password string `json:"password"`
}

func GetUserInfo(ctx *fiber.Ctx) error {
	return response.OkWithData(fiber.Map{
		"userName": "Soybean",
		"roles":    []string{"R_SUPER"},
		"avatar":   "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif",
		"userId":   0,
		"buttons":  []string{"B_CODE1", "B_CODE2", "B_CODE3"},
	}, ctx)
}

func GetUserList(ctx *fiber.Ctx) error {
	//  获取请求参数
	current := ctx.QueryInt("current", 1)
	size := ctx.QueryInt("size", 10)
	username := ctx.Query("username", "")
	// 查询用户列表
	result, err := services.GetUserList(username, current, size)
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
	result, err := services.CreateUser(userReq.Username, userReq.Password)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	return response.OkWithData(result, ctx)
}
