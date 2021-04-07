package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"time"
)

type Login struct {
	Username string `uri:"username" validate:"checkUsername"`
}

func checkUsername(f validator.FieldLevel) bool {
	return f.Field().String() != "admin"
}

type User struct {
	Name     string    `form:"name"json:"name" uri:"name" binding:"required"`
	Age      int       `form:"age" json:"age" uri:"age" binding:"required,gt=10"`
	Birthday time.Time `form:"birthday" json:"birthday" uri:"birthday" time_format:"2006-01-01" time_utc:"1"`
}

// 参数校验
func main() {
	g := gin.Default()
	validate := validator.New()
	g.GET("/", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindQuery(&user); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": 0, "data": user})
	})

	g.GET("/:username", func(c *gin.Context) {
		if err := validate.RegisterValidation("checkUsername", checkUsername); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"code": 0, "error": err.Error()})
			return
		}

		var login Login
		if err := c.ShouldBindUri(&login); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"code": 0, "error": err.Error()})
			return
		}

		if err := validate.Struct(login); err != nil {
			for _, e := range err.(validator.ValidationErrors) {
				c.JSON(http.StatusBadRequest, gin.H{"code": 0, "error": e})
			}
			return
		}

		c.JSON(http.StatusOK, gin.H{"code": 0, "data": login})
	})

	g.Run(":80")
}


