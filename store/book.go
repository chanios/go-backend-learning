package store

import (
	"backend/model"
	"errors"

	"gorm.io/gorm"
)

type BookStore struct {
	db *gorm.DB
}

func NewBookStore(db *gorm.DB) *BookStore {
	return &BookStore{
		db: db,
	}
}

func (bs *BookStore) List(limit int) ([]model.Book, int64, error) {
	var (
		books []model.Book
		count int64
	)
	bs.db.Model(&books).Count(&count)
	bs.db.
		Limit(limit).
		Order("created_at desc").Find(&books)
	return books, count, nil
}

func (bs *BookStore) GetByID(id uint) (*model.Book, error) {
	var b model.Book
	if err := bs.db.First(&b, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
	}
	return &b, nil
}

func (bs *BookStore) Delete(b *model.Book) error {
	return bs.db.Delete(b).Error
}

func (bs *BookStore) Update(b *model.Book) error {
	tx := bs.db.Begin()
	if err := tx.Model(b).Updates(b).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (bs *BookStore) Create(b *model.Book) error {
	tx := bs.db.Begin()
	if err := tx.Create(&b).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
