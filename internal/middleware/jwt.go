package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

var jwtSecret []byte

func init() {
	secret, err := os.ReadFile("/run/secrets/jwt_secret")
	if err != nil {
		log.Fatalf("Ошибка чтения JWT Secret: %v", err)
	}
	jwtSecret = secret
	fmt.Println("JWT Secret загружен успешно") // Только для теста
}

func JWTAuthInterceptor(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/Singup", "/Login":
			next.ServeHTTP(w, r)
			return
		}

		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			log.Println("there are no strings in the authorization header")
			return
		}

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

		next.ServeHTTP(w, r)
	})
}
