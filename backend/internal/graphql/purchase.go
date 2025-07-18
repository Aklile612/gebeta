package graphql

import (
	"context"
	"log"

	"github.com/aklile/recipe-backend/internal/config"
	hasura "github.com/machinebox/graphql"
)


type RecipePurchaseInfo struct {
	ID          string
	IsPaid      bool
	Price       float64
	OwnerUserID string
}

func InsertRecipePurchase(userID, recipeID string, amount float64) error{
	adminSecret:= config.LoadADMINSecret()

	req:= hasura.NewRequest(`
		mutation($user_id: uuid!, $recipe_id: uuid,$amount: numeric!){
			insert_recipe_purchases_one(object:{
				user_id: $user_id,
				recipe_id: $recipe_id,
				amount: $amount,
			}){
				id	
			}
		}
	`)
	req.Var("user_id",userID)
	req.Var("recipe_id",recipeID)
	req.Var("amount",amount)

	req.Header.Set("x-hasura-admin-secret", string(adminSecret))
	var resp struct{
		InsertRecipePurchasesOne struct{
			ID string `json:"id"`
		}`json:"insert_recipe_purchases_one"`
	}

	err:= Client.Run(context.Background(),req,&resp)

	if err != nil{
		return err
	}
	log.Println("Added to the recipe purshase table")
	return nil
}

func GetRecipePurchaseInfo(userID,requesterUserID string) (*RecipePurchaseInfo,error){

	
}