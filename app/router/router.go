package router

import (
	"github.com/gofiber/contrib/swagger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"soybean-admin-go/app/controller"
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/middleware"
)

func Initalize(router *fiber.App) {

	// allow cors
	router.Use(cors.New())

	// swagger

	api := router.Group("/api")

	api.Use(swagger.New(swagger.Config{
		BasePath: "/api",
		FilePath: "./docs/swagger.json",
		Path:     "docs",
	}))

	// Init Casbin for Role-based Authorization Control (RBAC)
	enforcer := database.Casbin()

	api.Use(middleware.Json)

	auth := api.Group("auth")

	auth.Post("login", controller.Login)

	// 需要认证的用户
	auth.Use(middleware.JwtMiddleware())

	auth.Post("refreshToken", controller.RefreshToken)

	auth.Get("getUserInfo", controller.GetUserInfo)

	user := api.Group("user")

	user.Post("create", controller.CreateUser)

	system := api.Group("systemManage")
	system.Use(middleware.JwtMiddleware())
	system.Get("getUserList", middleware.AuthorizeCasbin(enforcer), controller.GetUserList)
	system.Get("getRoleList", middleware.AuthorizeCasbin(enforcer), controller.GetRoleList)
	system.Get("getAllRoles", middleware.AuthorizeCasbin(enforcer), controller.GetAllRoles)
	system.Get("/getMenuList/v2", middleware.AuthorizeCasbin(enforcer), controller.GetMenuList)
	system.Get("getAllPages", middleware.AuthorizeCasbin(enforcer), controller.GetAllPages)
	system.Post("role", middleware.AuthorizeCasbin(enforcer), controller.AddRole)
	system.Put("role/:id", middleware.AuthorizeCasbin(enforcer), controller.EditRole)

	//api.Post("/user", controller.CreateUser)
	//api.Get("/users", middleware.JwtMiddleware(), controller.GetUser)
	//api.Post("/login", controller.UserLogin)
	//
	//needAuth := api.Group("/menu")
	//
	//// Menu API
	//needAuth.Use(middleware.JwtMiddleware())

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})

}
