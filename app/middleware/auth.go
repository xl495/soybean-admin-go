package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"os"
	"soybean-admin-go/app/utils/response"
	"strings"
	"time"
)

func JwtMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		//	获取请求头中的token
		tokenString := c.Get("Authorization")

		if tokenString == "" {
			return response.FailWithUnauthorized(nil, "未登录", c)
		}

		bearerToken := strings.Split(tokenString, " ") // Bearer token

		if len(bearerToken) != 2 {
			return response.FailWithUnauthorized(nil, "未登录", c)
		}

		token := bearerToken[1]

		parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			return response.FailWithUnauthorized(nil, "Token 解析失败", c)
		}

		if parseToken != nil {
			if parseToken.Claims == nil || parseToken.Claims.(jwt.MapClaims)["userId"] == nil {
				return response.FailWithUnauthorized(nil, "未登录", c)
			}

			expireAt := time.Unix(int64(parseToken.Claims.(jwt.MapClaims)["exp"].(float64)), 0)

			if time.Now().Unix() > expireAt.Unix() {
				return response.FailWithUnauthorized(nil, "登录失效", c)
			}

			claims := parseToken.Claims.(jwt.MapClaims)

			c.Locals("user", claims)

			c.Locals("userId", claims["userId"])

		} else {
			return response.FailWithUnauthorized(nil, "未登录", c)
		}

		return c.Next()
	}
}
