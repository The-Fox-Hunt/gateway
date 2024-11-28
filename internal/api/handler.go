package api

import (
	"encoding/json"
	"github.com/The-Fox-Hunt/gateway/internal/model"
	"io"
	"log"
	"net/http"
)

type Handler struct {
	aS AuthService
}

func NewHandler(authS AuthService) *Handler {
	return &Handler{aS: authS}
}

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	jsn, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var data model.SignupData
	err = json.Unmarshal(jsn, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	resp, err := h.aS.SignUp(r.Context(), data)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsnResp, err := json.Marshal(resp)

	w.WriteHeader(http.StatusCreated)
	w.Write(jsnResp)
}
