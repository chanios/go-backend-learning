package handler

import "backend/book"

type Handler struct {
	bookStore book.Store
	validator *Validator
}

func NewHandler(bs book.Store) *Handler {
	v := NewValidator()
	return &Handler{
		bookStore: bs,
		validator: v,
	}
}
