package model

type UserRole struct {
	GormBase
	//	角色名称
	RoleName string `json:"roleName"`
	//	角色编码
	RoleCode string `json:"roleCode"`
	//	角色描述
	RoleDesc string `json:"roleDesc"`
	//	状态
	Status string `json:"status"`
}
