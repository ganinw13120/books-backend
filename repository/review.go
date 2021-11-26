package repository

import (
	"gorm.io/gorm"

	"books-backend/model"
)

type reviewRepository struct {
	db *gorm.DB
}

type IReviewRepository interface {
	GetReviewList() ([]model.Review, error)
}

func NewReviewRepository(db *gorm.DB) reviewRepository {
	return reviewRepository{db: db}
}

func (b reviewRepository) GetReviewList() ([]model.Review, error) {
	books := make([]model.Review, 0)
	err := b.db.Table(model.Tables.Review).Find(&books).Error
	return books, err
}
