package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golangApi.test/controllers"
	"golangApi.test/models"
)

func main() {

	r := gin.Default()

	//Models

	db := models.Setup()

	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Go Api"})
	})

	r.GET("/products", controllers.GetProducts)
	r.POST("/products", controllers.CreateProduct)

	r.Run(":8080")

}
