package main

import (
	"github.com/The-Fox-Hunt/gateway/internal/clients/auth"
	"net/http"

	"github.com/The-Fox-Hunt/gateway/internal/api"
)

func main() {

	authClient := auth.NewClient()

	handler := api.NewHandler(authClient)

	// Привязываем маршрут "/" к функции handleRoot
	http.HandleFunc("/signup", handler.HandleSignUp)

	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
