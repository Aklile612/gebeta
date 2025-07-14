package graphql

import (
	

	"github.com/aklile/recipe-backend/internal/config"
	hasura "github.com/machinebox/graphql"
)


var Client *hasura.Client
var endPOINT = config.ENDPoint()
func InitClient(){
	endPoint:= endPOINT
	Client= hasura.NewClient(string(endPoint))
}