package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func MiddleWare(c *gin.Context) {
	fmt.Println("1.开始中间件")
	startTime := time.Now()
	c.Set("name", "zhongming")

	// 调用具体执行函数
	c.Next()

	status := c.Writer.Status()
	fmt.Println("2.结束中间件：", status)
	endTime := time.Now()
	fmt.Printf("3.中间件耗时：%d秒\n", endTime.Second()-startTime.Second())
}

// 中间件练习
func main() {
	g := gin.Default()

	// 全局中间件
	//g.Use(MiddleWare)
	{
		g.GET("/mid", MiddleWare, func(c *gin.Context) {
			name, _ := c.Get("name")
			fmt.Println("4.name:", name)

			c.JSON(http.StatusOK, gin.H{"name": name})
		})
	}

	g.Run(":80")
}
