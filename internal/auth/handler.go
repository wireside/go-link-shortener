package auth

import (
	"encoding/json"
	"errors"
	"io"
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
		var payload LoginRequest
		err := json.NewDecoder(req.Body).Decode(&payload)
		if err != nil {
			if errors.Is(err, io.EOF) {
				httputil.BadRequest(w, "Request body is empty, JSON payload is required")
				return
			}

			httputil.BadRequest(w, err.Error())
			return
		}

		if payload.Email == "" {
			httputil.BadRequest(w, "Email is required but missing")
			return
		}
		if payload.Password == "" {
			httputil.BadRequest(w, "Password is required but missing")
			return
		}

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
