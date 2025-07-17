package handlers

type PurchaseEventPayload struct {
	Event struct {
		Op   string `json:"op"`
		Data struct {
			New struct {
				ID          string `json:"id"`
				UserID      string `json:"user_id"`
				RecipeID    string `json:"recipe_id"`
				PurchasedAt string `json:"purchased_at"`
				Amount      string `json:"amount"`
			} `json:"new"`
		} `json:"data"`
	} `json:"event"`
}


