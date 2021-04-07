package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	// 创建路由
	g := gin.Default()
	// 绑定路由规则，执行函数
	// gin.Context封装了request、response
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "get hello gin")
	})

	// 获取api参数
	// 可以通过Context的Param方法来获取API参数
	// http://localhost/钟鸣/30
	g.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		action = strings.Trim(action, "/")
		c.String(http.StatusOK, "用户名：%s,方法名：%s", name, action)
	})

	// 获取url参数
	// URL参数可以通过DefaultQuery()或Query()方法获取
	// DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
	// API ? name=zs
	g.GET("/user", func(c *gin.Context) {
		name := c.DefaultQuery("name", "钟鸣")
		age, _ := strconv.Atoi(c.DefaultQuery("age", "0"))
		c.String(http.StatusOK, "用户名：%s,年龄：%d", name, age)
	})

	// 获取表单（post）数据
	// 表单传输为post请求，http常见的传输格式为四种：
	// application/json
	// application/x-www-form-urlencoded
	// application/xml
	// multipart/form-data
	// 表单参数可以通过PostForm()方法获取，该方法默认解析的是x-www-form-urlencoded或from-data格式的参数
	g.POST("/user", func(c *gin.Context) {
		types := c.DefaultPostForm("type", "post")
		name := c.DefaultPostForm("name", "钟鸣")
		age, _ := strconv.Atoi(c.DefaultPostForm("age", "30"))
		c.String(http.StatusOK, "接口请求方式：%s,用户名：%s,年龄：%d", types, name, age)
	})

	g.Run(":80")
}
