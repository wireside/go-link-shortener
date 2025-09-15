package request

import (
	"encoding/json"
	"errors"
	"io"
)

var ErrNilBody = errors.New("request body is nil")

func Decode[T any](body io.ReadCloser) (*T, error) {
	if body == nil {
		return nil, ErrNilBody
	}
	defer body.Close()

	var payload T
	err := json.NewDecoder(body).Decode(&payload)
	if err != nil {
		return nil, err
	}

	return &payload, nil
}
