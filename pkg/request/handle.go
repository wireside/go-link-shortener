package request

import (
	"errors"
	"io"
	"net/http"

	"go-adv-demo/pkg/response"
)

func HandleBody[T any](w http.ResponseWriter, r *http.Request) (*T, error) {
	body, err := Decode[T](r.Body)
	if err != nil {
		if errors.Is(err, io.EOF) {
			response.BadRequest(w, "request body is empty, JSON payload is required")
			return nil, err
		}

		response.BadRequest(w, err.Error())
		return nil, err
	}

	err = Validate[*T](body)
	if err != nil {
		response.BadRequest(w, err.Error())
		return nil, err
	}

	return body, nil
}
