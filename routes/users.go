package routes

import (
	"fmt"
	"net/http"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

func signUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := user.Save(); err != nil {
		fmt.Println("Error saving user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user_id": user.ID})
}