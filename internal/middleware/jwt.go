package middleware

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/The-Fox-Hunt/gateway/config"
	"github.com/The-Fox-Hunt/gateway/internal/model"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte

func init() {
	jwtSecret, err := config.GetSecret("JWT_SECRET")
	if err != nil {
		log.Fatalf("Ошибка загрузки JWT: %v", err)
	}

	fmt.Println("JWT загружен:", jwtSecret)
}

func JWTAuthInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("there are no strings in the authorization header")
			return
		}

		log.Printf("Received raw token: '%s'", token)

		token = strings.TrimPrefix(token, "Bearer ")
		token = strings.TrimSpace(token)

		parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("incorrect token signing method")
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			log.Printf("Failed to parse token: %s", err)
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
			return
		}

		if !parsedToken.Valid {
			log.Printf("Ivalid token")
			http.Error(w, "Unauthorized: missing token", http.StatusUnauthorized)
		}

		claims, ok := parsedToken.Claims.(jwt.MapClaims)
		if !ok {
			log.Println("Failed to extract claims from token")
			http.Error(w, "Unauthorized: invalid token claims", http.StatusUnauthorized)
			return
		}

		username, ok := claims["username"].(string)
		if !ok || username == "" {
			log.Println("Username not found in token claims")
			http.Error(w, "Unauthorized: username missing", http.StatusUnauthorized)
			return
		}

		log.Printf("Get username from token: '%s'", username)

		//  Добавляем `username` в `context`
		ctx := context.WithValue(r.Context(), model.Username, username)
		log.Printf("Authenticated user: %s", username)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
