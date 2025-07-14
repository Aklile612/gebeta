package main

import (
	"fmt"
	"log"

	"github.com/aklile/recipe-backend/internal/config"
	"github.com/aklile/recipe-backend/internal/graphql"
	"github.com/aklile/recipe-backend/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	graphql.InitClient(string(config.ENDPoint()))
	fmt.Println("GraphQL client initialized:", graphql.Client != nil)

	router:= gin.Default()

	router.POST("/register",handlers.RegisterHandler)

	err:= router.Run(":8081")

	if err!=nil{
		log.Fatal("Failed to run the server",err)
	}
}