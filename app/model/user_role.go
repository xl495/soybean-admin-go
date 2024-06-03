package model

type UserRole struct {
	GormBase
	// 用户ID
	UserID uint `json:"user_id"`
	// 角色ID
	RoleID uint `json:"role_id"`
	//	角色名称
	RoleName string `json:"role_name"`
	//	角色编码
	RoleCode string `json:"role_code"`
	//	状态
	Status int `json:"status"`
}
