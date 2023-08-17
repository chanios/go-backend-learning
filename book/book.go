package book

import "backend/model"

type Store interface {
	List(limit int) ([]model.Book, int64, error)
	GetByID(id uint) (*model.Book, error)
	// GetByName(string) (*model.Book, error)
	Create(*model.Book) error
	Update(*model.Book) error
	Delete(*model.Book) error
}
