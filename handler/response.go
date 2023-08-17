package handler

import "backend/model"

type bookResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type bookListResponse struct {
	Books      []*bookResponse `json:"books"`
	BooksCount int64           `json:"booksCount"`
}

func newBookResponse(b *model.Book) *bookResponse {
	r := new(bookResponse)
	r.ID = b.ID
	r.Name = b.Name
	r.Description = b.Description

	return r
}

func newBookListResponse(books []model.Book, count int64) *bookListResponse {
	r := new(bookListResponse)
	r.Books = make([]*bookResponse, 0)

	for _, b := range books {
		br := new(bookResponse)
		br.Name = b.Name
		br.ID = b.ID
		br.Description = b.Description
		r.Books = append(r.Books, br)
	}
	r.BooksCount = count
	return r
}
