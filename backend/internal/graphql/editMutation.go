package graphql

import (
	"context"

	"github.com/aklile/recipe-backend/internal/config"
	hasura "github.com/machinebox/graphql"
)

func CheckRecipeOwnership(userID,recipeId string)(bool,error){

	adminSecret:= config.LoadADMINSecret()

	req:= hasura.NewRequest(`
		query($user_id: uuid!, $recipe_id: uuid!) {
			recipes_by_pk(id: $recipe_id) {
				user_id
			}
		}
	`)
	req.Var("user_id",userID)
	req.Var("recipe_id",recipeId)
	req.Header.Set("x-hasura-admin-secret",string(adminSecret))


	var resp struct {
		RecipesByPk *struct {
			UserID string `json:"user_id"`
		} `json:"recipes_by_pk"`
	}


	err:= Client.Run(context.Background(),req,&resp)

	if err!= nil || resp.RecipesByPk==nil{
		return false,err
	}

	return resp.RecipesByPk.UserID == userID,nil

}

func UpdateRecipe(id, title, description, imageURL, difficulty string, prepTime, cookTime int, categoryID string, isPaid bool, price float64)error{
	
}