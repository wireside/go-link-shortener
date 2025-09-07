package auth

import (
	"fmt"
	"net/http"

	"go-adv-demo/configs"
)

type AuthHandler struct{}

func NewAuthHandler(router *http.ServeMux, conf *configs.Config) {
	handler := &AuthHandler{}

	router.HandleFunc("/auth/login", handler.login())
	router.HandleFunc("/auth/register", handler.register())
}

func (handler AuthHandler) login() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("login")
	}
}

func (handler AuthHandler) register() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("register")
	}
}
