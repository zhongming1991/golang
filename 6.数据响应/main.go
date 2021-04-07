package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"net/http"
)

func main() {
	g := gin.Default()

	// json
	g.GET("/json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": "json ok"})
	})

	// 结构体
	g.GET("/struct", func(c *gin.Context) {
		var data Data
		data.Username = "struct"
		data.Password = "123456"

		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": data})
	})

	// xml
	g.GET("/xml", func(c *gin.Context) {
		var data Data
		data.Username = "xml"
		data.Password = "123456"
		c.XML(http.StatusOK, gin.H{"code": http.StatusOK, "data": data})
	})

	// yaml
	g.GET("/yaml", func(c *gin.Context) {
		var data Data
		data.Username = "yaml"
		data.Password = "123456"
		c.YAML(http.StatusOK, gin.H{"code": http.StatusOK, "data": data})
	})

	// protobuf
	g.GET("/protobuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		// 定义数据
		label := "label"
		// 传protobuf格式数据
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		c.ProtoBuf(http.StatusOK, data)
	})

	g.Run(":80")
}

type Data struct {
	Username string
	Password string
}
