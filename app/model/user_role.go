package model

type UserRole struct {
	GormBase
	//	角色名称
	RoleName string `json:"roleName" gorm:"comment: 角色名称"`
	//	角色编码
	RoleCode string `json:"roleCode" gorm:"comment: 角色编码"`
	//	角色描述
	RoleDesc string `json:"roleDesc" gorm:"comment: 角色描述"`
	//	状态
	Status string `json:"status"  gorm:"default:1;comment: 状态 1 正常 2 禁用"`
}
