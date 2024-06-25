package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/config"
	"github.com/nugrohoarr/golang-mux-jwt-boilerplate/helper"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")

		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
			helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
			return
		}

		tokenStr := c.Value
		claims := &config.JWTClaims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})
		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			if v.Errors == jwt.ValidationErrorExpired {
				response := map[string]string{"message": "Unauthorized, token expired"}
				helper.ResponseJSON(w, http.StatusUnauthorized, response)
				return
			}
			helper.ResponseJSON(w, http.StatusBadRequest, err.Error())
			return
		}
		if !tkn.Valid {
			response := map[string]string{"message": "Unauthorized, token invalid"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
