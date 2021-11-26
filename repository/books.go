package repository

import (
	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

type IBookRepository interface {
}

func NewBookRepository(db *gorm.DB) bookRepository {
	return bookRepository{db: db}
}
