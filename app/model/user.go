package model

import (
	"database/sql/driver"
	"encoding/json"
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
	UserRoles UserRoleType `gorm:"type:longtext"`
	//	状态
	Status string `json:"status" gorm:"default:1; comment: 状态 1 正常 2 禁用"`
}

type UserRoleType []string

// BeforeSave 自定义方法，用于在保存记录前将 UserRoles 转换为 JSON 字符串
//func (u *User) BeforeSave(_ *gorm.DB) (err error) {
//	if len(u.UserRoles) > 0 {
//		rolesJson, err := json.Marshal(u.UserRoles)
//		if err != nil {
//			return err
//		}
//		u.Roles = string(rolesJson)
//	}
//	return nil
//}

// AfterFind 自定义方法，用于在查找记录后将 JSON 字符串转换为 UserRoles
//func (u *User) AfterFind(tx *gorm.DB) (err error) {
//	if u.Roles != "" {
//		err = json.Unmarshal([]byte(u.Roles), &u.UserRoles)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

// Scan 解码json字符串
func (role *UserRoleType) Scan(val interface{}) error {
	b, _ := val.([]byte)
	return json.Unmarshal(b, role)
}

// Value 编码json
func (role UserRoleType) Value() (value driver.Value, err error) {
	return json.Marshal(role)
}
