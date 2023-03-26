package main

import (
	"flikko/config"
	"flikko/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	config.Connect()
	routes.UserRoute(router)
	routes.ProductRoute(router)
	router.Run(":8080")
}
