package auth

import (
	"context"
	"net/http"
	"strings"
)

type key int

const UserKey key =0
func JWTMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		authHeader:=r.Header.Get("Authorization")
		if authHeader == ""{
			http.Error(w,"missing authorization header",http.StatusUnauthorized)
			return 
		}
		parts:= strings.Split(authHeader," ")
		if len(parts)!= 2 || parts[0] != "Bearer"{
			http.Error(w,"invalid authorization header",http.StatusUnauthorized)
			return
		}

		tokenStr:= parts[1]
		claims,err:= ValidateJWT(tokenStr)
		if err!= nil{
			http.Error(w,"invalid token",http.StatusUnauthorized)
		}

		ctx:= context.WithValue(r.Context(),UserKey,claims.UserID)
		next.ServeHTTP(w,r.WithContext(ctx))
	})
}