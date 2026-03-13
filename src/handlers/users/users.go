package users

import (
	"net/http"
	"step/service"
)

type UsersHandler struct {
	service *service.Service
	mux     *http.ServeMux
}

func NewUsersHandler(service *service.Service, mux *http.ServeMux) *UsersHandler {
	return &UsersHandler{
		service: service,
		mux:     mux,
	}
}
