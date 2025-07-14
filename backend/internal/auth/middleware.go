package auth

import "net/http"


func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandleFunc(func(w http.))
}