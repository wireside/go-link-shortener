package link

import (
	"net/http"

	"go-adv-demo/configs"
)

type LinkHandler struct {
	*LinkHandlerDeps
}

type LinkHandlerDeps struct {
	Config *configs.Config
}

func NewLinkHandler(router *http.ServeMux, deps *LinkHandlerDeps) *LinkHandler {
	handler := &LinkHandler{
		deps,
	}

	router.HandleFunc("POST /link", handler.create())
	router.HandleFunc("GET /{alias}", handler.goTo())
	router.HandleFunc("PATCH /link/{id}", handler.update())
	router.HandleFunc("DELETE /link/{id}", handler.delete())

	return handler
}

func (handler *LinkHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (handler *LinkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (handler *LinkHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}

func (handler *LinkHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
