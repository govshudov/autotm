package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"

	"auth/internal/services"
)

type HTTPHandler struct {
	service *services.CartService
}

func NewHTTPHandler(service *services.CartService) *HTTPHandler {
	return &HTTPHandler{service: service}
}

func (h *HTTPHandler) RegisterRoutes(r chi.Router) {
	r.Method("GET", "/ds", http.HandlerFunc(h.addItem))
}

func (h *HTTPHandler) addItem(w http.ResponseWriter, r *http.Request) {
	fmt.Println("addItem")
	err := h.service.Test(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	ints, err := w.Write([]byte("OK"))
	if err != nil {
		fmt.Println("Error writing response", err.Error())
	}
	fmt.Println(ints)
	return
}
