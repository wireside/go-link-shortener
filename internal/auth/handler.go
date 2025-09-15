package auth

import (
	"net/http"

	"go-adv-demo/configs"
	"go-adv-demo/pkg/request"
	"go-adv-demo/pkg/response"
)

type AuthHandler struct {
	*AuthHandlerDeps
}

type AuthHandlerDeps struct {
	Config *configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	handler := &AuthHandler{
		deps,
	}

	router.HandleFunc("POST /auth/login", handler.login())
	router.HandleFunc("POST /auth/register", handler.register())
}

func (handler *AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		_, err := request.HandleBody[LoginRequest](w, req)
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
	return func(w http.ResponseWriter, req *http.Request) {
		body, err := request.HandleBody[RegisterRequest](w, req)
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
