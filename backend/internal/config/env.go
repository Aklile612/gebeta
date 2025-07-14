package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadEnv(){
	err:= godotenv.Load(".env")

	if err != nil {
		log.Println("NO env file found or error loading it",err)
	}
}

func LoadJWTSecret()[]byte{
	LoadEnv()

	return []byte(os.Getenv("HASURA_GRAPHQL_JWT_SECRET"))
}
func LoadADMINSecret()[]byte{
	LoadEnv()

	return []byte(os.Getenv("HASURA_GRAPHQL_ADMIN_SECRET"))
}

func ENDPoint() []byte{
	LoadEnv()
	return []byte(os.Getenv("END_POINT"))

}