package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/someGet/:name", getting1)
	r.GET("/someGet", getting2)
	r.POST("/somePost", postting)
	r.POST("/someGetPost", gettingPostting)
	r.POST("/someGetPostMap", gettingAndPostMap)

	v1 := r.Group("/v1")
	{
		v1.GET("/read", readEndPoint)
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/login", loginEndPoint)
	}

	uploadFile := r.Group("/uploadFile")
	{
		uploadFile.POST("/singleFile1", uploadSingleFile1)
		uploadFile.POST("/singleFile2", uploadSingleFile2)
		uploadFile.POST("/multiFiles", uploadMultiFiles)
	}

	r.Run(":8081")
}

func getting1(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func getting2(c *gin.Context) {
	firstName := c.DefaultQuery("firstName", "Guest")
	lastName := c.Query("lastName")

	c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
}

func postting(c *gin.Context) {
	message := c.PostForm("message")
	result := c.DefaultPostForm("result", "nothing")

	c.JSON(http.StatusOK, gin.H{
		"status":  "posted",
		"message": message,
		"result":  result,
	})
}

func gettingPostting(c *gin.Context) {
	id := c.Query("id")
	page := c.Query("page")
	name := c.PostForm("name")
	message := c.PostForm("message")

	c.String(http.StatusOK, "id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}

func gettingAndPostMap(c *gin.Context) {
	ids := c.QueryMap("ids")
	names := c.PostFormMap("names")

	fmt.Printf("ids: %v; names: %v", ids, names)
	c.JSON(http.StatusOK, gin.H{
		"ids":   ids,
		"names": names,
	})
}

func readEndPoint(c *gin.Context) {
	c.String(http.StatusOK, "read")
}

func loginEndPoint(c *gin.Context) {
	c.String(http.StatusOK, "login")
}

func uploadSingleFile1(c *gin.Context) {
	file, handler, err := c.Request.FormFile("file")
	filename := handler.Filename
	log.Println("Received file:", filename)

	out, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.String(http.StatusOK, "Upload file success")
}

func uploadSingleFile2(c *gin.Context) {
	file, _ := c.FormFile("file")
	log.Println("Received file:", file.Filename)
	c.SaveUploadedFile(file, file.Filename)

	c.String(http.StatusOK, "Upload file success")
}

func uploadMultiFiles(c *gin.Context) {
	form, _ := c.MultipartForm()
	files := form.File["files"]

	for _, file := range files {
		log.Println(file.Filename)
		c.SaveUploadedFile(file, file.Filename)
	}

	c.String(http.StatusOK, "Upload multiFile success")
}
