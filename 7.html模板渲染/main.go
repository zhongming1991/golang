package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.LoadHTMLGlob("html/**/*")
	g.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index/index.html", gin.H{"title": "第一个go Html", "address": "湖南省长沙市"})
	})
	g.Run(":80")
}
