package model

import (
	"encoding/json"
	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	GormBase
	UserName string `json:"userName"`
	// 昵称
	NickName string `json:"nickName" gorm:"default: null; comment: 昵称"`
	// 性别
	UserGender string `json:"userGender" gorm:"default: 1; comment: 性别 1 男 2 女 3 未知"`
	// 手机号
	UserPhone string `json:"userPhone" gorm:"unique;default: null; comment: 手机号码"`
	// 邮箱
	UserEmail string `json:"userEmail" gorm:"unique;default: null; comment: 邮箱"`
	// 密码
	Password string `json:"password"`
	// 头像
	Avatar string `json:"avatar" gorm:"default: null; comment: 头像"`
	// 角色
	UserRoles []string `json:"userRoles" gorm:"-"`
	// 用于存储 UserRoles 的 JSON 字符串
	Roles string `json:"-"`
	//	状态
	Status string `json:"status" gorm:"default:1; comment: 状态 1 正常 2 禁用"`
}

// BeforeSave 自定义方法，用于在保存记录前将 UserRoles 转换为 JSON 字符串
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	if len(u.UserRoles) > 0 {
		rolesJson, err := json.Marshal(u.UserRoles)
		if err != nil {
			return err
		}
		u.Roles = string(rolesJson)
	}
	return nil
}

// AfterFind 自定义方法，用于在查找记录后将 JSON 字符串转换为 UserRoles
func (u *User) AfterFind(_ *gorm.DB) (err error) {
	if u.Roles != "" {
		err = json.Unmarshal([]byte(u.Roles), &u.UserRoles)
		if err != nil {
			return err
		}
	}
	return nil
}
