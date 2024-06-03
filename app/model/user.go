package model

// User 用户模型
type User struct {
	GormBase
	Username string `json:"username"`
	// 昵称
	Nickname string `json:"nickname"`
	// 性别
	Gender string `json:"gender"`
	// 手机号
	Phone string `json:"phone;unique"`
	// 邮箱
	Email string `json:"email;unique"`
	// 密码
	Password string `json:"password"`
	// 头像
	Avatar string `json:"avatar"`
	// 角色
	Role string `json:"role"`
}
