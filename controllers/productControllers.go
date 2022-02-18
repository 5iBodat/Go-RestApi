package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golangApi.test/models"
	"gorm.io/gorm"
)

type ProductInput struct {
	Id    int     `json:"id" gorm:"primaryKey"`
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

// get product by code

func GetProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Products
	db.First(&product, c.Param("id"))
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// create product
func CreateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses create product
	product := models.Products{
		Code:  input.Code,
		Name:  input.Name,
		Price: input.Price,
	}
	db.Create(&product)
	c.JSON(http.StatusOK, gin.H{"data": product})
}

// update product

func UpdateProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var input ProductInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	//proses update product
	var product models.Products
	db.First(&product, c.Param("id"))
	product.Code = input.Code
	product.Name = input.Name
	product.Price = input.Price
	db.Save(&product)

	c.JSON(http.StatusOK, gin.H{"data": product})

}

// delete product

func DeleteProduct(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)
	var product models.Products
	db.First(&product, c.Param("id"))
	db.Delete(&product)
	c.JSON(http.StatusOK, gin.H{"data": "Product deleted"})
}
