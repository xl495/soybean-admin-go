package response

import (
	"github.com/gofiber/fiber/v2"
)

// Response 定义了标准化API响应的结构体。
type Response struct {
	Code string      `json:"code" description:"HTTP状态码或应用特定错误码"`
	Data interface{} `json:"data" description:"响应数据，可以是任何JSON可序列化的类型"`
	Msg  string      `json:"msg" description:"描述操作结果的消息"`
}

const (
	ERROR   = "0007" // 应用特定错误码
	SUCCESS = "0000" // 应用特定成功码

	UNAUTHORIZED = "401" //	未登录
)

// Result 根据提供的参数构建并发送一个标准化的API响应。
// @summary 构建并返回一个标准化的API响应
// @returns error - 如果发送响应时出现问题，则返回错误
func Result(code string, data interface{}, msg string, c *fiber.Ctx) error {
	// 开始时间（如果需要）
	return c.JSON(Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

// Ok 发送一个成功响应，其中数据字段为空，并带有默认消息。
// @summary 发送一个成功响应，无数据且使用默认消息
// @returns error - 如果发送响应时出现问题，则返回错误
func Ok(c *fiber.Ctx) error {
	return Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

// OkWithMessage 发送一个成功响应，其中数据字段为空，但带有自定义消息。
// @summary 发送一个成功响应，无数据但使用自定义消息
// @returns error - 如果发送响应时出现问题，则返回错误
func OkWithMessage(message string, c *fiber.Ctx) error {
	return Result(SUCCESS, map[string]interface{}{}, message, c)
}

// OkWithData 发送一个成功响应，其中包含提供的数据和默认消息。
// @summary 发送一个成功响应，包含数据和默认消息
// @returns error - 如果发送响应时出现问题，则返回错误
func OkWithData(data interface{}, c *fiber.Ctx) error {
	return Result(SUCCESS, data, "操作成功", c)
}

// OkWithDetailed 发送一个成功响应，其中包含提供的数据和自定义消息。
// @summary 发送一个成功响应，包含数据和自定义消息
// @returns error - 如果发送响应时出现问题，则返回错误
func OkWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
	return Result(SUCCESS, data, message, c)
}

// Fail 发送一个失败的响应，其中数据字段为空，并带有默认消息。
// @summary 发送一个失败的响应，无数据且使用默认消息
// @returns error - 如果发送响应时出现问题，则返回错误
func Fail(c *fiber.Ctx) error {
	return Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

// FailWithMessage 发送一个失败的响应，其中数据字段为空，但带有自定义消息。
// @summary 发送一个失败的响应，无数据但使用自定义消息
// @returns error - 如果发送响应时出现问题，则返回错误
func FailWithMessage(message string, c *fiber.Ctx) error {
	return Result(ERROR, map[string]interface{}{}, message, c)
}

// FailWithDetailed 发送一个失败的响应，其中包含提供的数据和自定义消息。
// @summary 发送一个失败的响应，包含数据和自定义消息
// @returns error - 如果发送响应时出现问题，则返回错误
func FailWithDetailed(data interface{}, message string, c *fiber.Ctx) error {
	return Result(ERROR, data, message, c)
}

// FailWithUnauthorized 发送一个失败的响应，其中包含提供的数据和自定义消息。
// @summary 发送一个失败的响应，包含数据和自定义消息
// @returns error - 如果发送响应时出现问题，则返回错误
func FailWithUnauthorized(data interface{}, message string, c *fiber.Ctx) error {
	return Result(UNAUTHORIZED, data, message, c)
}
