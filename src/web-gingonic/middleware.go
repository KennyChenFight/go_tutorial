package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	// 使用日誌
	router.Use(gin.Logger())
	// 使用Panic處理方案
	router.Use(gin.Recovery())
	// 路由經由多個函數
	router.GET("/index", index(), chain)

	authorized := router.Group("/author")
	authorized.Use(authRequired())
	{
		authorized.POST("/login", loginEndPoint)
	}

	router.Run(":8081")
}

func index() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("before middleware")
	}
}

func chain(c *gin.Context) {
	fmt.Println("index chainFun")
}

func loginEndPoint(c *gin.Context) {
	value, exist := c.Get("request")

	fmt.Println("authorized loginEndPoint --exist:", exist, "--value", value)
}

func authRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("brfore middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after middleware")
	}
}
