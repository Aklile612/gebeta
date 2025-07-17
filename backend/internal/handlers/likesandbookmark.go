package handlers

import (
	"net/http"

	"github.com/aklile/recipe-backend/internal/graphql"
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
	err:= graphql.InesrtRecipeLikes(userID,recipeID)

	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"cannot like the recipe","detail":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"the recipe liked succefully"})

}


func ADDRecipeBookmarks(c *gin.Context){
	recipeID:= c.Param("id")

	userIDInterf,exists:= c.Get("user_id")

	if !exists{
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
		return
	}

	userID ,ok := userIDInterf.(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user id not valied"})
		return
	}

	err:= graphql.InsertUserRecipeBookmark(userID,recipeID)
	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"cannot like the recipe","detail":err.Error()})
		return
	}

	c.JSON(http.StatusOK,gin.H{"message":"the recipe liked succefully"})

}
