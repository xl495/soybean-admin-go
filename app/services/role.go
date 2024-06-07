package services

import (
	"soybean-admin-go/app/database"
	"soybean-admin-go/app/model"
	"soybean-admin-go/app/utils/response"
)

func GetRoleList(roleName, roleCode, status string, current int, size int) (response.PageResult, error) {
	var userRoles []model.UserRole
	var total int64

	// 开始构造查询
	query := database.DB.Find(&userRoles)

	if roleName != "" {
		query = query.Where("role_name LIKE ?", "%"+roleName+"%")
	}

	if roleCode != "" {
		query = query.Where("role_code LIKE ?", "%"+roleCode+"%")
	}

	if status != "" {
		query = query.Where("status = ?", status)
	}

	query.Count(&total)

	query.Limit(size).Offset((current - 1) * size).Find(&userRoles)

	// 确保返回的 users 不是 nil
	if userRoles == nil {
		userRoles = []model.UserRole{}
	}

	return response.PageResult{
		Records: userRoles,
		Total:   total,
		Current: current,
		Size:    size,
	}, nil
}

func GetAllRoles() ([]model.UserRole, error) {
	var userRoles []model.UserRole
	query := database.DB.Find(&userRoles)

	if query.Error != nil {
		return userRoles, query.Error
	}
	return userRoles, nil
}

func AddRole(roleName, roleCode, roleDesc, status string) (model.UserRole, error) {
	var userRole model.UserRole
	userRole.RoleName = roleName
	userRole.RoleCode = roleCode
	userRole.RoleDesc = roleDesc
	userRole.Status = status
	query := database.DB.Create(&userRole)
	if query.Error != nil {
		return userRole, query.Error
	}
	return userRole, nil
}

func EditRole(id int, roleName, roleCode, roleDesc, status string) (model.UserRole, error) {
	var userRole model.UserRole
	userRole.RoleName = roleName
	userRole.RoleCode = roleCode
	userRole.RoleDesc = roleDesc
	userRole.Status = status

	query := database.DB.Model(&userRole).Where("id = ?", id).Updates(&userRole)

	if query.Error != nil {
		return userRole, query.Error
	}

	return userRole, nil
}

func RemoveRoles(ids []int) error {
	query := database.DB.Where("id IN ?", ids).Delete(&model.UserRole{})

	if query.Error != nil {
		return query.Error
	}
	return nil
}
