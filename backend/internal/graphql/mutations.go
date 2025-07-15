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
	func InsertRecipe(title,desc,imageURL, difficulty string, prepTime, cookTime int, userID, categoryID string, isPaid bool, price float64)(models.Recipe,error){

		var adminSecret = config.LoadADMINSecret()
		req:= hasura.NewRequest(`
			mutation($title: String!, $description: String!, $featured_image: String!, $difficulty: String!, $prep_time_minutes: Int!, $cook_time_minutes: Int!, $user_id: uuid!, $category_id: uuid!, $is_paid: Boolean!, $price: numeric!){
				insert_recipes_one(object:{
					title: $title,
					description: $description,
					featured_image: $featured_image,
					difficulty: $difficulty,
					prep_time_minutes: $prep_time_minutes,
					cook_time_minutes: $cook_time_minutes,
					user_id: $user_id,
					category_id: $category_id,
					is_paid: $is_paid,
					price: $price
				}){
					id
					title
					featured_image	
				}
			}
		`)
		req.Var("title", title)
		req.Var("description", desc)
		req.Var("featured_image", imageURL)
		req.Var("difficulty", difficulty)
		req.Var("prep_time_minutes", prepTime)
		req.Var("cook_time_minutes", cookTime)
		req.Var("user_id", userID)
		req.Var("category_id", categoryID)
		req.Var("is_paid", isPaid)
		req.Var("price", price)
		req.Header.Set("x-hasura-admin-secret", string(adminSecret))

		var resp struct{
			InsertRecipe models.Recipe `json:"insert_recipes_one"`
		}

		err := Client.Run(context.Background(),req,&resp)

		return resp.InsertRecipe,err
	}