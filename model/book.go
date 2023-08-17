package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	AuthorID    uint
	Name        string
	Description string
}
