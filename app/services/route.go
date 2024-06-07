package services

import (
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
)

func GetConstantRoutes() []model.Menu {
	var menus []model.Menu

	database.DB.Where("parent_id = 0").Find(&menus)

	return menus
}
