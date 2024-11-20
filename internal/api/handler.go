package api

import (
	"fmt"
	"net/http"

	"github.com/The-Fox-Hunt/gateway/internal/model"
)

type Handler struct {
	Field1 string
	field2 string
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Handle(first string, second string) string {
	data := model.RequestData{
		Field1: first,
		Field2: second,
	}
	// send to service layer
	data.Result = first + second
	return data.Result
}

func (h *Handler) HandleRoot(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		// Устанавливаем статус ответа 200 OK
		w.WriteHeader(http.StatusCreated)
		// Отправляем тело ответ
		fmt.Fprintln(w, "Hello, world!")
	} else {
		// Если метод не GET, возвращаем 405 Method Not Allowed
		http.Error(w, "Только GET поддерживается!", http.StatusMethodNotAllowed)
	}
}
