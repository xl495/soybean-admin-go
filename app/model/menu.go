package model

// Menu 菜单模型模型
type Menu struct {
	GormBase
	// 父级ID
	ParentId uint `json:"parentId" gorm:"default:0;comment: 父级id"`
	// 菜单类型
	MenuType string `json:"menuType" gorm:"default:1;comment: 菜单类型"`
	// 菜单名称
	MenuName string `json:"menuName" gorm:"comment: 菜单名称"`
	// 路由名称
	RouteName string `json:"routeName" gorm:"comment: 路由名称"`
	// 路由路径
	RoutePath string `json:"routePath" gorm:"comment: 路由路径"`
	// 组件路径
	Component string `json:"component" gorm:"comment: 组件路径"`
	// 排序
	Order int `json:"order" gorm:"comment: 排序"`
	// 多语言 key
	I18NKey string `json:"i18nKey" gorm:"comment: 多语言 key"`
	// icon
	Icon string `json:"icon" gorm:"comment: 图标"`
	// icon 类型
	IconType string `json:"iconType" gorm:"comment: 图标类型"`
	// 是否启用
	Status string `json:"status" gorm:"default:1;"`
}
