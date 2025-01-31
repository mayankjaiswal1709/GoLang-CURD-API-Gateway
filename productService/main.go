package main

import (
	controllers "productService/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/product", controllers.AddProduct)
	router.DELETE("/deleteProduct", controllers.DeleteProduct)
	router.DELETE("/deleteProduct/:id", controllers.DeleteProductByID)
	router.GET("/getProduct", controllers.GetProduct)
	router.GET("/getProduct/:id", controllers.GetProductByID)
	router.PUT("/updateProduct/:id", controllers.UpdateProduct)


	router.Run(":8082") 
}
