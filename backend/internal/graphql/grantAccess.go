package graphql

import (
	"context"
	"log"
	"time"

	"github.com/aklile/recipe-backend/internal/config"
	hasura "github.com/machinebox/graphql"
)


func GrantAccess(userID, recipeID string) error{
	adminSecret:= config.LoadADMINSecret()

	req:= hasura.NewRequest(`
		mutation ($user_id: uuid!, $recipe_id: uuid!, $granted_at: timestamptz!){
			insert_recipe_access_one(object:{
				user_id: $user_id,
				recipe_id: $recipe_id,
				granted_at: $granted_at,
			},on_conflict: {
				contraint: recipe_access_user_id_recipe_id_key,
				update_columns: [access_granted_at]  
			}){
				id	
			}
		}
	`)

	req.Var("user_id",userID)
	req.Var("recipe_id",recipeID)
	req.Var("access_granted_at",time.Now().UTC().Format(time.RFC3339))

	req.Header.Set("x-hasura-admin-secret", string(adminSecret))

	var resp struct{

		InsertRecipeAccessOne struct{
			ID   string `json:"id"`
		}`json:"insert_recipe_access_one"`
	}


	err:= Client.Run(context.Background(),req,&resp)

	if err != nil{
		return err
	}

	log.Printf("Granted access: user %s -> recipe %s\n", userID, recipeID)
	return nil

}