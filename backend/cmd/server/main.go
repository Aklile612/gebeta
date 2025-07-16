package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aklile/recipe-backend/internal/auth"
	"github.com/aklile/recipe-backend/internal/config"
	"github.com/aklile/recipe-backend/internal/graphql"
	"github.com/aklile/recipe-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	
	config.LoadEnv()
	fmt.Println("END_POINT from os.Getenv:", os.Getenv("END_POINT"))
	endpoint := string(config.ENDPoint())
	fmt.Println("Loaded END_POINT from config.ENDPoint():", endpoint)

	if endpoint == "" {
		log.Fatal("❌ END_POINT is empty. Please check your .env and config.")
	}

	graphql.InitClient(endpoint)
	if graphql.Client == nil {
		log.Fatal("❌ GraphQL client is still nil after InitClient")
	} else {
		fmt.Println("✅ GraphQL client initialized correctly")
	}
	// go build -o servercmd && ./servercmd 
	router := gin.Default()

	router.POST("/register", handlers.RegisterHandler)
	router.POST("/login",handlers.LoginHandler)

	authGroup:= router.Group("/")
	authGroup.Use(auth.JWTMiddleware())
	authGroup.POST("/add_recipes",handlers.AddRecipeHandler)

	err := router.Run(":8081")

	if err != nil {
		log.Fatal("Failed to run the server", err)
	}
}