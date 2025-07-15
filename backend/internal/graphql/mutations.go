	package graphql

	import (
		"context"
		"fmt"
		

		"github.com/aklile/recipe-backend/internal/config"
		
		"github.com/aklile/recipe-backend/internal/models"
		hasura "github.com/machinebox/graphql"
	)
	func InsertUser(name, email, password string) (models.User, error) {
		fmt.Printf("InsertUser called. Client: %v\n", Client)
		
		if Client == nil {
			return models.User{}, fmt.Errorf("graphql client not initialized")
		}
		var adminSecret = config.LoadADMINSecret()
		
		fmt.Printf("Admin Secret: %s\n", string(adminSecret))

		req := hasura.NewRequest(`
			mutation($full_name:String!,$email:String!,$password:String!){
				insert_users_one(object:{
				full_name:$full_name,
				email:$email,
				password:$password
				}){
				id
				email
				full_name
				}
			}
		`)
		req.Var("full_name",name)
		req.Var("email",email)
		req.Var("password",password)
		req.Header.Set("x-hasura-admin-secret",string(adminSecret))

		var resp struct{
			InsertUser models.User  `json:"insert_users_one"`

		}
		err:= Client.Run(context.Background(),req,&resp)
		if err != nil {
			fmt.Printf("❌ InsertUser GraphQL error: %v\n", err)
			return models.User{}, fmt.Errorf("insert user failed: %w", err)
		}

		return resp.InsertUser, nil

	}
