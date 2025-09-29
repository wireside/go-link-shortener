package link

import (
	"net/http"

	"go-adv-demo/pkg/request"
	"go-adv-demo/pkg/response"
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
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		createdLink, err := handler.linkRepository.Create(link)
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		response.Created(w, createdLink)
	}
}

func (handler *LinkHandler) goTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")

		link, err := handler.linkRepository.GetByHash(hash)
		if err != nil {
			response.NotFound(w, err)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (handler *LinkHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
