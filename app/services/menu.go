package services

import (
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils/response"
)

func GetMenuTreeList(current int, size int) (response.PageResult, error) {
	var menus []model.Menu
	var total int64

	// 开始构造查询
	query := database.DB

	// 使用分页和获取用户列表
	query.Offset((current - 1) * size).Limit(size).Find(&menus)

	// 计算总数
	query.Model(&model.Menu{}).Count(&total)

	// 确保返回的 users 不是 nil
	if menus == nil {
		menus = []model.Menu{}
	}

	return response.PageResult{
		Records: menus,
		Total:   total,
		Current: current,
		Size:    size,
	}, nil
}
