package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRecipeLikes(c *gin.Context){
	recipeID := c.Param("id")

	userIDInterf, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}
	userID, ok := userIDInterf.(string)

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not valied"})
		return
	}
	

}
