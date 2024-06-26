package controller

import (
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/services"
	"soybean-admin-go/app/utils/response"
)

type LogReq struct {
	// 用户名
	// required: true
	// example: admin
	UserName string `json:"userName"`
	// 密码
	// required: true
	// example: 123456
	Password string `json:"password"`
}

func RefreshToken(ctx *fiber.Ctx) error {
	return response.OkWithMessage("refreshToken", ctx)
}

func Login(ctx *fiber.Ctx) error {

	// 获取请求参数
	var logReq LogReq
	// 解析参数
	if err := ctx.BodyParser(&logReq); err != nil {
		return response.FailWithMessage("参数解析错误", ctx)
	}

	// 登录
	result, err := services.Login(logReq.UserName, logReq.Password)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	return response.OkWithData(result, ctx)
}
