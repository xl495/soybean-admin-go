package model

// Menu 菜单模型模型
type Menu struct {
	GormBase
	// 父级ID
	ParentId uint `json:"parentId"`
	// 菜单类型
	MenuType string `json:"menuType"`
	// 菜单名称
	MenuName string `json:"menuName"`
	// 路由名称
	RouteName string `json:"routeName"`
	// 路由路径
	RoutePath string `json:"routePath"`
	// 组件路径
	Component string `json:"component"`
	// 排序
	Order int `json:"order"`
	// 多语言 key
	I18NKey string `json:"i18nKey"`
	// icon
	Icon string `json:"icon"`
	// icon 类型
	IconType string `json:"iconType"`
	// 是否启用
	Status string `json:"status" gorm:"default:1;"`
}
