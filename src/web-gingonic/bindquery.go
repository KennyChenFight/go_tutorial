package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name    string `form:"name" json:"name"`
	Address string `form:"address" json:"address"`
}

func main() {
	router := gin.Default()
	router.Any("/testing", startPage)

	router.Run(":8081")
}

func startPage(c *gin.Context) {
	var person Person
	if c.BindQuery(&person) == nil {
		log.Println("======= Only Bind Query String =======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	if c.Bind(&person) == nil {
		log.Println("======= Only Bind Query String and Post Data =======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	if c.BindJSON(&person) == nil {
		log.Println("======= Only Bind JSON =======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	c.String(http.StatusOK, "Success")
}
