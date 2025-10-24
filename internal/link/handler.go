package link

import (
	"net/http"
	"strconv"

	"go-adv-demo/internal/auth"
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

	router.Handle("POST /link", handler.create())
	router.Handle("GET /{hash}", handler.goTo())
	router.Handle("PATCH /link/{id}", auth.IsAuthed(handler.update()))
	router.Handle("DELETE /link/{id}", handler.delete())

	return handler
}

func (handler *LinkHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkCreateRequest](w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)
		for {
			existedLink, _ := handler.linkRepository.GetByHash(link.Hash)
			if existedLink == nil {
				break
			}

			link.GenerateHash()
		}

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
			response.NotFound(w, err.Error())
			return
		}

		http.Redirect(w, r, link.Url, http.StatusTemporaryRedirect)
	}
}

func (handler *LinkHandler) update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := request.HandleBody[LinkUpdateRequest](w, r)
		if err != nil {
			return
		}
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		existedLink, err := handler.linkRepository.GetByID(uint(id))
		if err != nil {
			response.NotFound(w, err.Error())
			return
		}

		existedLink.Url = body.Url
		existedLink.Hash = body.Hash

		link, err := handler.linkRepository.Update(existedLink)
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		response.OK(w, link)
	}
}

func (handler *LinkHandler) delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			response.BadRequest(w, err.Error())
			return
		}

		existedLink, err := handler.linkRepository.GetByID(uint(id))
		if err != nil {
			response.NotFound(w, err.Error())
			return
		}

		err = handler.linkRepository.Delete(existedLink.ID)
		if err != nil {
			response.InternalServerError(w, err.Error())
			return
		}

		response.OK(w, nil)
	}
}
