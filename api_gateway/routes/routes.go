package routes

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	userService := "http://localhost:8081" 

	router.POST("/user", func(c *gin.Context) {
		c.Redirect(307, userService+"/user")
	})

	router.DELETE("/deleteUser", func(c *gin.Context) {
		c.Redirect(307, userService+"/deleteUser")
	})

	router.DELETE("/deleteUser/:id", func(c *gin.Context) {
		c.Redirect(307, userService+"/deleteUser/"+c.Param("id"))
	})

	router.GET("/getUser", func(c *gin.Context) {
		c.Redirect(307, userService+"/getUser")
	})

	router.GET("/getUser/:id", func(c *gin.Context) {
		c.Redirect(307, userService+"/getUser/"+c.Param("id"))
	})

	router.PUT("/updateUser/:id", func(c *gin.Context) {
		c.Redirect(307, userService+"/updateUser/"+c.Param("id"))
	})
}

func ProductRoutes(router *gin.Engine) {
	productService := "http://localhost:8082" 

	router.POST("/product", func(c *gin.Context) {
		c.Redirect(307, productService+"/product")
	})

	router.DELETE("/product/:id", func(c *gin.Context) {
		c.Redirect(307, productService+"/product/"+c.Param("id"))
	})

	router.GET("/getProduct", func(c *gin.Context) {
		c.Redirect(307, productService+"/getProduct")
	})

	router.GET("/getProduct/:id", func(c *gin.Context) {
		c.Redirect(307, productService+"/getProduct/"+c.Param("id"))
	})

	router.PUT("/updateProduct/:id", func(c *gin.Context) {
		c.Redirect(307, productService+"/updateProduct/"+c.Param("id"))
	})
}
