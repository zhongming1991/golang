package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建路由
	g := gin.Default()

	// 单个文件上传
	// multipart/form-data格式用于文件上传
	// gin文件上传与原生的net/http方法类似，不同在于gin把原生的request封装到c.Request中
	g.POST("/upload", uploadFile)

	// 多个文件上传
	g.POST("uploadFiles", uploadFiles)

	// 启动web服务，监听80端口
	// Run()不写端口默认为8080
	g.Run(":80")
}

// 上传单个文件
func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	var errMsg string
	if err != nil {
		errMsg = err.Error()
		goto resp
	}

	if file.Header.Get("Content-Type") != "image/png" {
		errMsg = "只允许上传png图片"
		goto resp
	}

	if file.Size > 1024 {
		errMsg = "上传图片不能大于1KB"
		goto resp
	}

	c.SaveUploadedFile(file, "D:/"+file.Filename)
	c.String(http.StatusOK, "文件上传成功")

resp:
	c.String(http.StatusInternalServerError, errMsg)
}

// 上传多个文件
func uploadFiles(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	// 获取上传文件列表
	files := form.File["files"]
	if len(files) == 0 {
		c.String(http.StatusBadRequest, "未获取到上传文件")
		return
	}

	// 循环保存上传的文件
	for index, file := range files {
		if err := c.SaveUploadedFile(file, "D:/"+file.Filename); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("第%d个文件[%s]上传失败：%s", index+1, file.Filename, err.Error()))
			return
		}
	}

	c.String(http.StatusOK, "文件上传成功")
}
