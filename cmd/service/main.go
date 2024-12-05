package main

import (
	"net/http"

	"github.com/The-Fox-Hunt/gateway/internal/clients/auth"
	"github.com/The-Fox-Hunt/gateway/internal/service"

	"github.com/The-Fox-Hunt/gateway/internal/api"
)

func main() {

	authClient := auth.NewClient()

	authService := service.NewService(authClient)

	handler := api.NewHandler(authService)

	// Привязываем маршрут "/" к функции handleRoot
	http.HandleFunc("/signup", handler.HandleSignUp)
	http.HandleFunc("/login", handler.HandleSignIn)

	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
