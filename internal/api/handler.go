package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/The-Fox-Hunt/gateway/internal/model"
)

type Handler struct {
	aS AuthService
}

func NewHandler(authS AuthService) *Handler {
	return &Handler{aS: authS}
}

func (h *Handler) HandleRequest(w http.ResponseWriter, r *http.Request, data interface{}, action func() (interface{}, error)) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	jsn, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Printf("failed to write response: %v", err)
		}
		return
	}

	err = json.Unmarshal(jsn, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		if _, err = w.Write([]byte(err.Error())); err != nil {
			log.Printf("failed to write response: %v", err)
		}
		return
	}

	log.Printf("Received data: %+v\n", data)

	resp, err := action()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("failed to write header: %v", err)
		return
	}

	jsnResp, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	if _, writeErr := w.Write(jsnResp); writeErr != nil {
		log.Printf("failed to write response body: %v", writeErr)
	}
}

func (h *Handler) HandleSignUp(w http.ResponseWriter, r *http.Request) {
	var data model.SignupData
	h.HandleRequest(w, r, &data, func() (interface{}, error) {
		return h.aS.SignUp(r.Context(), data)
	})
}

func (h *Handler) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	var data model.SignInData
	h.HandleRequest(w, r, &data, func() (interface{}, error) {
		return h.aS.SignIn(r.Context(), data)
	})
}

func (h *Handler) HandleChangePassword(w http.ResponseWriter, r *http.Request) {
	var data model.ChangePasswordData

	h.HandleRequest(w, r, &data, func() (interface{}, error) {
		return h.aS.ChangePassword(r.Context(), data)
	})
}
