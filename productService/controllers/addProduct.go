package controllers

import (
	"log"
	"net/http"
	"productService/db"
	"productService/models"

	"github.com/gin-gonic/gin"
)

func AddProduct(c *gin.Context) {
	var product models.Product

	// Bind incoming JSON to the user model
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := db.GetCollection("productData")

	insertResult, err := collection.InsertOne(c.Request.Context(), product)
	if err != nil {
		log.Printf("Failed to insert document: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert document"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully", "id": insertResult.InsertedID})
}
