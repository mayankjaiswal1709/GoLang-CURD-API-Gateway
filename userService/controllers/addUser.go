package controller

import (
	"userService/db"
	"userService/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("userData")

	insertResult, err := collection.InsertOne(c.Request.Context(), user)
	if err != nil {
		log.Printf("Failed to insert document: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully", "id": insertResult.InsertedID})
}
