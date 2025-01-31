package main

import (
	"api_gateway/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.UserRoutes(router)
	routes.ProductRoutes(router)

	router.Run(":8070")
}
