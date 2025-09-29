package auth

import (
	"net/http"

	"go-adv-demo/pkg/request"
	"go-adv-demo/pkg/response"
)

type AuthHandler struct {
}

func NewAuthHandler(router *http.ServeMux) *AuthHandler {
	handler := &AuthHandler{}

	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())

	return handler
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LoginRequest](w, r)
		if err != nil {
			return
		}

		res := LoginResponse{
			Token: "123",
		}

		response.OK(w, res)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[RegisterRequest](w, r)
		if err != nil {
			return
		}

		res := RegisterResponse{
			Name:  body.Name,
			Email: body.Email,
			Token: "123",
		}

		response.OK(w, res)
	}
}
