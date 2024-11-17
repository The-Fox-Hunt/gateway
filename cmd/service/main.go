package main

import (
	"net/http"

	"github.com/The-Fox-Hunt/gateway/internal/api"
)

func main() {

	handler := api.NewHandler()

	// Привязываем маршрут "/" к функции handleRoot
	http.HandleFunc("/", handler.HandleRoot)

	// Запускаем сервер
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
