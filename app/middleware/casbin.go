package middleware

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils/response"
)

func AuthorizeCasbin(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user model.User

		userId := c.Locals("userId").(float64)

		// 查找当前用户

		if userId == 0 {
			return response.FailWithMessage("查找当前用户失败!", c)
		}

		database.DB.Find(&user, userId)

		var role = user.UserRoles[0]

		if role == "" {
			return response.FailWithMessage("当前用户未设置角色!", c)
		}

		err := e.LoadPolicy()
		if err != nil {
			return response.FailWithMessage("加载配置文件失败!", c)
		}

		accepted, err := e.Enforce(fmt.Sprint(role), c.OriginalURL(), c.Method()) // id - url - method || 1 - /api/admin/users - GET

		if err != nil {
			return response.FailWithMessage("授权失败!", c)
		}

		if !accepted {
			return response.FailWithMessage("未授权!", c)
		}
		return c.Next()
	}
}
