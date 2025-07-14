package handlers

import (
	"encoding/json"
	"net/http"
)


type LoginRequest struct{
	Email string  `json:"emial"`
	Password string `json:"password"`
}

type LoginResponse struct{

	Token  string `json:"token"`
	Error string `json:"error,omitempty"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request){
	var req LoginRequest
	err:= json.NewDecoder(r.Body).Decode(&req)

	if err!= nil{
		http.Error(w,"Bad request",http.StatusBadRequest)
		return
	}
	

}