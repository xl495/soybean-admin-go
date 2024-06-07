package controller

import (
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/services"
	"soybean-admin-go/app/utils/response"
)

type MenuTreeResp struct {
	model.Menu
	Children []MenuTreeResp `json:"children"`
	Name     string         `json:"name"`
	Path     string         `json:"path"`
}
type MenuTreeListResp struct {
	Pid      uint               `json:"pid"`
	Label    string             `json:"label"`
	ID       uint               `json:"id"`
	Children []MenuTreeListResp `json:"children"`
}

func SetResourceMenuTree(menuList []model.Menu) []MenuTreeResp {
	var menuTreeList []MenuTreeResp
	for _, menu := range menuList {
		menuTree := MenuTreeResp{
			Menu: menu,
			Name: menu.MenuName,
			Path: menu.RouteName,
		}
		if menu.ParentId == 0 {
			if menuTree.Children == nil {
				menuTree.Children = []MenuTreeResp{}
			}
			menuTreeList = append(menuTreeList, menuTree)

			continue
		} else {
			for index, item := range menuTreeList {
				if item.ID == menu.ParentId {
					if menuTree.Children == nil {
						menuTree.Children = []MenuTreeResp{}
					}
					menuTreeList[index].Children = append(menuTreeList[index].Children, menuTree)
				}
			}
		}
	}
	return menuTreeList
}

func SetResourceMenuTreeList(menuList []model.Menu) []MenuTreeListResp {
	var menuTreeList []MenuTreeListResp
	for _, menu := range menuList {
		menuTree := MenuTreeListResp{
			Pid:   menu.ParentId,
			Label: menu.MenuName,
			ID:    menu.ID,
		}
		if menu.ParentId == 0 {
			menuTreeList = append(menuTreeList, menuTree)
			continue
		} else {
			for index, item := range menuTreeList {
				if item.ID == menu.ParentId {
					menuTreeList[index].Children = append(menuTreeList[index].Children, menuTree)
				}
			}
		}
	}
	return menuTreeList
}

func GetMenuList(ctx *fiber.Ctx) error {
	current := ctx.QueryInt("current", 1)
	size := 500
	result, err := services.GetMenuTreeList(current, size)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	resultTree := SetResourceMenuTree(result.Records.([]model.Menu))

	return response.OkWithData(response.PageResult{
		Records: resultTree,
		Total:   result.Total,
		Current: result.Current,
		Size:    result.Size,
	}, ctx)
}

func GetMenuTreeList(ctx *fiber.Ctx) error {
	result, err := services.GetMenuTreeList(1, 500)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	resultTree := SetResourceMenuTreeList(result.Records.([]model.Menu))

	return response.OkWithData(resultTree, ctx)
}

// GetAllPages 获取所有菜单 name
func GetAllPages(ctx *fiber.Ctx) error {
	current := ctx.QueryInt("current", 1)
	size := 500
	result, err := services.GetMenuTreeList(current, size)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	var resultPageName []string

	for _, v := range result.Records.([]model.Menu) {
		strings := append(resultPageName, v.RouteName)
		resultPageName = strings
	}

	return response.OkWithData(resultPageName, ctx)
}
