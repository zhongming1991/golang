package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Cookie的缺点
// 1、不安全，明文
// 2、增加带宽消耗
// 3、可以被禁用
// 4、cookie有上限

// 模拟实现权限验证中间件
// 有2个路由，login和index
// login用于设置cookie
// index是访问查看信息的请求
// 在请求index之前，先跑中间件代码，检验是否存在cookie
func main() {
	g := gin.Default()

	// 登录
	g.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		password := c.Query("password")

		if !(username == "admin" && password == "123456") {
			c.JSON(http.StatusOK, gin.H{"code": 0, "data": "用户名或者密码错误"})
			return
		}

		_, err := c.Cookie("isLogin")
		if err != nil {
			// 给客户端设置cookie
			// maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			// secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("isLogin", "true", 60*120, "/", "localhost", false, true)
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "登录成功"})
	})

	g.GET("/index", validIsLogin, func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": "欢迎来到首页"})
	})

	g.Run(":80")
}

func validIsLogin(c *gin.Context) {

	if isLogin, err := c.Cookie("isLogin"); err == nil {
		if b, _ := strconv.ParseBool(isLogin); b {
			startTime := time.Now()

			// 执行调用主函数
			c.Next()

			endTime := time.Now()
			log.Printf("接口耗时：%d秒", endTime.Second()-startTime.Second())
			return
		}
	}

	// 未授权
	c.JSON(http.StatusUnauthorized, gin.H{"code": 0, "data": "请先登录"})
	// 若验证不通过，不再调用后续的函数处理
	c.Abort()
}