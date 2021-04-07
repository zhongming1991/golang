package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
)

// 自定义验证v10
// Validator 是基于 tag（标记）实现结构体和单个字段的值验证库，它包含以下功能：
// 使用验证 tag（标记）或自定义验证器进行跨字段和跨结构体验证。
// 关于 slice、数组和 map，允许验证多维字段的任何或所有级别。
// 能够深入 map 键和值进行验证。
// 通过在验证之前确定接口的基础类型来处理类型接口。
// 处理自定义字段类型（如 sql 驱动程序 Valuer）。
// 别名验证标记，它允许将多个验证映射到单个标记，以便更轻松地定义结构体上的验证。
// 提取自定义的字段名称，例如，可以指定在验证时提取 JSON 名称，并在生成的 FieldError 中使用该名称。
// 提取自定义的字段名称，例如，可以指定在验证时提取 JSON 名称，并在生成的 FieldError 中使用该名称。
// 可自定义 i18n 错误消息。
// Web 框架 gin 的默认验证器。

// 标签
//通过以上章节的内容，读者应该已经了解到 Validator 是一个基于 tag（标签），实现结构体和单个字段的值验证库。
//标签	描述
//eq	等于
//gt	大于
//gte	大于等于
//lt	小于
//lte	小于等于
//ne	不等于
//max	最大值
//min	最小值
//oneof	其中一个
//required	必需的
//unique	唯一的
//isDefault	默认值
//len	长度
//email	邮箱格式
func main() {
	validate := validator.New()

	// 验证变量
	email := "admin#admin.com"
	err := validate.Var(email, "required,email")
	if err != nil {
		errors := err.(validator.ValidationErrors)
		fmt.Println("var验证：", errors)
	} else {
		fmt.Println("var验证：", email)
	}

	// 结构体验证
	var login struct {
		Username string `json:"username" validate:"required"`
		Age      int    `json:"age" validate:"required,gte=18,lte=60"`
		Email    string `json:"email" validate:"required,email"`
	}

	login.Username = "zhongming"
	login.Age = 1
	login.Email = "406224709@qq.com"

	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		return field.Tag.Get("json")
	})

	errs := validate.Struct(login)
	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			fmt.Println("结构体验证", err)
		}
	}
}
