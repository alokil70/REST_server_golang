package book

import "test_RESTserver_01/internal/author"

type Book struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Author []author.Author `json:"author"`
}
