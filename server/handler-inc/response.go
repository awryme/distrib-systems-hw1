package handler_inc

import (
	"encoding/json"
	"io"
)

type errResponse struct {
	Err   string `json:"err,omitempty"`
}

type okResponse struct {
	Query int    `json:"query",bson:"query"`
	Resp  int    `json:"resp",bson:"resp"`
}

func renderErr(writer io.Writer, err error) {
	res := errResponse{
		Err: err.Error(),
	}
	renderJson(writer, res)
}

func renderJson(writer io.Writer, v interface{}) {
	_ = json.NewEncoder(writer).Encode(v)
}