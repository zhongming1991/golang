package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()

	g.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello go")
	})

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, fmt.Sprintf("地址：%s,未找到", c.Request.RequestURI))
	})

	g.Run(":80")
}
