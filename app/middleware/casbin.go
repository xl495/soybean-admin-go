package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"soybean-admin-go/app/utils/response"
)

func AuthorizeCasbin(e *casbin.Enforcer) fiber.Handler {
	return func(c *fiber.Ctx) error {

		user := c.Locals("user").(jwt.MapClaims)

		if user["userId"].(float64) == 0 {
			return response.FailWithMessage("未登录", c)
		}

		roles := user["userRole"].([]interface{})

		if len(roles) == 0 {
			return response.FailWithMessage("当前用户未设置角色!", c)
		}

		// 可能存在多个权限
		err := e.LoadPolicy()

		if err != nil {
			return response.FailWithMessage("加载配置文件失败!", c)
		}

		// 判断策略中是否存在
		for _, v := range roles {
			isSuccess, err := e.Enforce(v, c.OriginalURL(), c.Method())
			if err != nil {
				return response.FailWithMessage("鉴权失败!", c)
			}
			if isSuccess {
				return c.Next()
			}
		}
		return response.FailWithMessage("未授权!", c)
	}
}
