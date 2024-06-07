package controller

import (
	"github.com/gofiber/fiber/v2"
	"soybean-admin-go/app/services"
	"soybean-admin-go/app/utils/response"
)

type AddRoleReq struct {
	// 角色名称
	// required: true
	// example: admin
	RoleName string `json:"roleName"`
	// 角色编码
	// required: true
	// example: 123456
	RoleCode string `json:"roleCode"`
	// 角色描述
	// required: false
	// example: 角色描述
	RoleDesc string `json:"roleDesc"`
	// 状态
	// required: true
	// example: 1
	Status string `json:"status"`
}

type RemoveRoleReq struct {
	Ids []int `json:"ids"`
}

func GetRoleList(ctx *fiber.Ctx) error {
	current := ctx.QueryInt("current", 1)
	size := ctx.QueryInt("size", 10)
	roleName := ctx.Query("roleName", "")
	roleCode := ctx.Query("roleCode", "")
	status := ctx.Query("status", "")
	// 查询用户列表
	result, err := services.GetRoleList(roleName, roleCode, status, current, size)
	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}
	return response.OkWithData(result, ctx)
}

func GetAllRoles(ctx *fiber.Ctx) error {
	result, err := services.GetAllRoles()
	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}
	return response.OkWithData(result, ctx)
}

func EditRole(ctx *fiber.Ctx) error {
	editId, err := ctx.ParamsInt("id")
	if err != nil {
		return response.FailWithMessage("参数错误", ctx)
	}
	var editRoleReq AddRoleReq
	if err := ctx.BodyParser(&editRoleReq); err != nil {
		return response.FailWithMessage("参数解析错误", ctx)
	}
	if editRoleReq.RoleName == "" || editRoleReq.RoleCode == "" || editRoleReq.Status == "" {
		return response.FailWithMessage("参数错误", ctx)
	}
	result, err := services.EditRole(editId, editRoleReq.RoleName, editRoleReq.RoleCode, editRoleReq.RoleDesc, editRoleReq.Status)
	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}
	return response.OkWithData(result, ctx)
}

func AddRole(ctx *fiber.Ctx) error {
	var addRoleReq AddRoleReq
	if err := ctx.BodyParser(&addRoleReq); err != nil {
		return response.FailWithMessage("参数解析错误", ctx)
	}
	if addRoleReq.RoleName == "" || addRoleReq.RoleCode == "" || addRoleReq.Status == "" {
		return response.FailWithMessage("参数错误", ctx)
	}
	result, err := services.AddRole(addRoleReq.RoleName, addRoleReq.RoleCode, addRoleReq.RoleDesc, addRoleReq.Status)
	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}
	return response.OkWithData(result, ctx)
}

func RemoveRoles(ctx *fiber.Ctx) error {
	var removeRoleReq RemoveRoleReq
	if err := ctx.BodyParser(&removeRoleReq); err != nil {
		return response.FailWithMessage("参数解析错误", ctx)
	}
	if len(removeRoleReq.Ids) == 0 {
		return response.FailWithMessage("参数错误", ctx)
	}

	err := services.RemoveRoles(removeRoleReq.Ids)

	if err != nil {
		return response.FailWithMessage(err.Error(), ctx)
	}

	return response.OkWithMessage("删除成功", ctx)
}
