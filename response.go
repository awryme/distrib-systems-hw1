package main

type Response struct {
	Query int    `json:"query",bson:"query"`
	Resp  int    `json:"resp",bson:"resp"`
	Err   string `json:"err,omitempty"`
}
