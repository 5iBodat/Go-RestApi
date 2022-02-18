package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golangApi.test/models"
	"gorm.io/gorm"
)

type ProductInput struct {
	Id    int     `json:"id" sql:"AUTO_INCREMENT" gorm:"default:1"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func GetProducts(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var products []models.Products
	db.Find(&products)
	c.JSON(http.StatusOK, gin.H{"data": products})
}

// create product
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input ProductInput
	c.BindJSON(&input)
	product := models.Products{
		Code:  input.Code,
		Name:  input.Name,
		Price: input.Price,
	}
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}
