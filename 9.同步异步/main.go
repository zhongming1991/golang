package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

// goroutine机制可以方便地实现异步处理
// 另外，在启动新的goroutine时，不应该使用原始上下文，必须使用它的只读副本
func main() {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	g := gin.Default()

	// 异步
	g.GET("/async", func(c *gin.Context) {
		context := c.Copy()

		// 异步执行
		go func() {
			time.Sleep(3 * time.Second)
			log.Printf("异步执行：%s\n", context.Request.URL.Path)
		}()

		c.String(http.StatusOK, "异步执行")
	})

	// 同步
	g.GET("/sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Printf("同步执行：%s\n", c.Request.URL.Path)
		c.String(http.StatusOK, "同步执行")
	})


	g.Run(":80")
}
