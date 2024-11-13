package api

import "github.com/The-Fox-Hunt/gateway/internal/model"

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
