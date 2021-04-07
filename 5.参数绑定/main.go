package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 登录参数结构体
type Login struct {
	Username string `form:"username" json:"username" xml:"username" uri:"username" binding:"required"`
	Password string `form:"password" json:"password" xml:"password" uri:"password" binding:"required"`
}

// 数据解析和绑定
func main() {
	g := gin.Default()

	// JSON数据绑定
	g.POST("/loginJson", func(c *gin.Context) {
		// 接收JSON参数
		var login Login
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": err.Error()})
			return
		}


		// 验证是否登录
		validLogin(login, c)
	})

	// form表单绑定
	g.POST("/loginForm", func(c *gin.Context) {
		var login Login
		if err := c.Bind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "-1", "error": err.Error()})
			return
		}

		// 验证是否登录
		validLogin(login, c)
	})

	// uri数据绑定
	g.GET("/loginUri/:username/:password", func(c *gin.Context) {
		var login Login
		if err := c.BindUri(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": "-1", "error": err.Error()})
			return
		}

		// 验证是否登录
		validLogin(login, c)
	})
	g.Run(":80")
}

// 公共验证登录方法
func validLogin(login Login, c *gin.Context)  {
	if login.Username != "zm" || login.Password != "123456" {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "error": "用户名或者密码错误"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "data": "登录成功"})
}
