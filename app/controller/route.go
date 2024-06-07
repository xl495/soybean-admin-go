package controller

import (
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/services"
	"soybean-admin-go/app/utils/response"
)

type Menu2Route struct {
	Name      string       `json:"name"`
	Path      string       `json:"path"`
	Component string       `json:"component"`
	Children  []Menu2Route `json:"children"`
}

func GetConstantRoutes(ctx *fiber.Ctx) error {
	var route []Menu2Route

	menus := services.GetConstantRoutes()

	for _, menu := range menus {
		route = append(route, Menu2Route{
			Name:      menu.MenuName,
			Path:      menu.RouteName,
			Component: menu.Component,
			Children:  []Menu2Route{},
		})
	}

	return response.OkWithData(route, ctx)
}
