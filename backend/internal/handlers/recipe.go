package handlers

import (
	"net/http"
	"strconv"

	"github.com/aklile/recipe-backend/internal/graphql"
	"github.com/aklile/recipe-backend/internal/media"
	"github.com/gin-gonic/gin"
)


func AddRecipeHandler(c *gin.Context){
	userIDInterf, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
		return
	}

	userID, ok := userIDInterf.(string)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}
	title:= c.PostForm("title")
	description := c.PostForm("description")
	prepTimeStr := c.PostForm("prep_time_minutes")
	cookTimeStr := c.PostForm("cook_time_minutes")
	difficulty := c.PostForm("difficulty")
	
	categoryName := c.PostForm("category_name")
	categoryID, err := graphql.GetOrCreateCatagoryID(categoryName)
if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get/create category"})
	return
}
	isPaid := c.PostForm("is_paid") == "true"
	priceStr := c.PostForm("price")

	file,fileHeader,err:= c.Request.FormFile("image")

	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"error uploading image"})
		return
	}
	defer file.Close()
	imageURL,err:= media.UploadImage(file,fileHeader)

	if err!= nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":"Failed to upload image"})
		return
	}

	prepTime,_:= strconv.Atoi(prepTimeStr)
	cookTime,_:= strconv.Atoi(cookTimeStr)
	price,_ := strconv.ParseFloat(priceStr,64)


	recipe,err:= graphql.InsertRecipe(title, description, imageURL, difficulty, prepTime, cookTime, userID, categoryID, isPaid, price)

	if err!= nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"Failed to create a recipe"})
		return
	}
	c.JSON(http.StatusOK,gin.H{"recipe":recipe})
}