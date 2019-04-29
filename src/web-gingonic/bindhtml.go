package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

type LoginForm struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", nil)
	})
	r.POST("/login", func(c *gin.Context) {
		var form LoginForm
		if c.ShouldBind(&form) == nil {
			if form.User == "kenny" && form.Password == "0000" {
				c.JSON(http.StatusOK, gin.H{
					"status": "you are logged in",
				})
			} else {
				c.JSON(401, gin.H{
					"status": "unauthorized",
				})
			}
		}
	})

	r.Run(":8081")
}

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "view.html", nil)
}

func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{
		"color": fakeForm.Colors,
	})
}
