package api

import (
	"net/http"
)

type Handler struct {
	aC AuthClient
}

func NewHandler(authC AuthClient) *Handler {
	return &Handler{aC: authC}
}

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	err := h.aC.DoSignUp(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
