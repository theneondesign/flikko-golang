package routes

import (
	"flikko/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.GET("/users/", controllers.GetUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.POST("/users/create", controllers.CreateUser)
	router.PUT("/users/update/:id", controllers.UpdateUser)
	router.DELETE("/users/delete/:id", controllers.DeleteUser)
}

func ProductRoute(router *gin.Engine) {
	router.GET("/products/", controllers.GetProducts)
	router.GET("/products/:id", controllers.GetProduct)
	router.POST("/products/create/", controllers.CreateProduct)
	router.PUT("/products/update/:id", controllers.UpdateProduct)
	router.DELETE("/products/delete/:id", controllers.DeleteProduct)
}
