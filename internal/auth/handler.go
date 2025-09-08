package auth

import (
	"fmt"
	"net/http"

	"go-adv-demo/configs"
	"go-adv-demo/pkg/httputil"
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
		fmt.Println(handler.Config.Auth.Secret)
		fmt.Println("login")

		res := LoginResponse{
			Token: "123",
		}

		httputil.OK(w, res)
	}
}

func (handler *AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
