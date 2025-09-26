package link

import (
	"net/http"
)

type LinkHandler struct {
	linkRepository *LinkRepository
}

func NewLinkHandler(router *http.ServeMux, linkRepository *LinkRepository) *LinkHandler {
	handler := &LinkHandler{
		linkRepository,
	}

	router.HandleFunc("POST /link", handler.create())
	router.HandleFunc("GET /{hash}", handler.goTo())
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
