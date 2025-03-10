package main

import (
	"log"
	"net/http"

	"github.com/The-Fox-Hunt/gateway/internal/api"
	"github.com/The-Fox-Hunt/gateway/internal/clients/auth"
	"github.com/The-Fox-Hunt/gateway/internal/middleware"
	"github.com/The-Fox-Hunt/gateway/internal/service"
)

func main() {

	authClient := auth.NewClient()

	authService := service.NewService(authClient)

	handler := api.NewHandler(authService)

	// Привязываем маршрут "/" к функции handleRoot
	http.Handle("/signup", http.HandlerFunc(handler.HandleSignUp))
	http.Handle("/login", http.HandlerFunc(handler.HandleSignIn))

	// Все остальные запросы проходят через JWTAuthInterceptor
	protectedRoutes := http.NewServeMux()
	protectedRoutes.Handle("/changepassword", http.HandlerFunc(handler.HandleChangePassword))
	//protectedRoutes.Handle("/logout", http.HandlerFunc(handler.HandleLogout))

	http.Handle("/", middleware.JWTAuthInterceptor(protectedRoutes))

	// Запускаем сервер
	log.Println("Server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
