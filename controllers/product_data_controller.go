package controllers

import (
	"flikko/config"
	"flikko/models"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	Products := []models.Product{}
	config.DB.Find(&Products)
	c.JSON(200, &Products)
}
func GetNearbyProducts(c *gin.Context) {
	Products := []models.Product{}
	config.DB.Find(&Products)
	c.JSON(200, &Products)
}

func GetProduct(c *gin.Context) {
	Product := models.Product{}
	config.DB.Where("id = ?", c.Param("id")).First(&Product)
	c.JSON(200, &Product)
}

func CreateProduct(c *gin.Context) {
	var Product models.Product
	c.BindJSON(&Product)
	config.DB.Create(&Product)
	c.JSON(200, &Product)
}

func UpdateProduct(c *gin.Context) {
	var Product models.Product
	config.DB.Where("id = ?", c.Param("id")).First(&Product)
	c.BindJSON(&Product)
	config.DB.Save(&Product)
	c.JSON(200, &Product)
}

func DeleteProduct(c *gin.Context) {
	var Product models.Product
	config.DB.Where("id = ?", c.Param("id")).Delete(&Product)
	c.JSON(200, &Product)
}
