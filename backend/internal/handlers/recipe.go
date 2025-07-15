package handlers

import (
	"net/http"
	"strconv"

	"github.com/aklile/recipe-backend/internal/media"
	"github.com/gin-gonic/gin"
)


func AddRecipeHandler(c *gin.Context){
	title:= c.PostForm("title")
	description := c.PostForm("description")
	prepTimeStr := c.PostForm("prep_time_minutes")
	cookTimeStr := c.PostForm("cook_time_minutes")
	difficulty := c.PostForm("difficulty")
	userID := c.PostForm("user_id")
	categoryID := c.PostForm("category_id")
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


	
}