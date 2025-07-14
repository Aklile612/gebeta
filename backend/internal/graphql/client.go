package graphql

import (
	"github.com/machinebox/graphql"
)

var Client *graphql.Client

func InitClient(endpoint string) {
	Client = graphql.NewClient(endpoint)
}
