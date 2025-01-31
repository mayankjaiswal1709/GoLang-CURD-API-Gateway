package main

import (
	controllers "userService/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/user", controllers.AddUser)
	router.DELETE("/deleteUser", controllers.DeleteUser)
	router.DELETE("/deleteUser/:id", controllers.DeleteUserByID)
	router.GET("/getUser", controllers.GetUser)
	router.GET("/getUser/:id", controllers.GetUserByID)
	router.PUT("/updateUser/:id", controllers.UpdateUser)

	router.Run(":8081") 
}
