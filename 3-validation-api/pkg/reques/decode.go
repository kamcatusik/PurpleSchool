package reques

import (
	"encoding/json"
	"io"
)

func Decode[T any](r io.Reader) (*T, error) {
	var emailLoad T

	err := json.NewDecoder(r).Decode(&emailLoad)
	if err != nil {
		return &emailLoad, err
	}
	return &emailLoad, nil
}
