package handler_inc

import (
	"encoding/json"
	"fmt"
	"io"
)

type requestObj struct {
	Query int `json:"query"`
}

func parseRequestObj(reader io.Reader) (*requestObj, error) {
	obj := &requestObj{}
	err := json.NewDecoder(reader).Decode(obj)
	if err != nil {
		return nil, fmt.Errorf("failed to parse request object: %w", err)
	}
	return obj, nil
}
