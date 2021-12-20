package repository

import (
	"gorm.io/gorm"

	"books-backend/model"
)

type reviewRepository struct {
	db *gorm.DB
}

type IReviewRepository interface {
	GetAllReviewList() ([]model.Review, error)
	GetReviewByBookName(name string) ([]model.Review, error)
	GetReviewByID(id int) (model.Review, error)
}

func NewReviewRepository(db *gorm.DB) reviewRepository {
	return reviewRepository{db: db}
}

func (repo reviewRepository) GetAllReviewList() ([]model.Review, error) {
	reviews := make([]model.Review, 0)
	err := repo.db.Table(model.Tables.Review).Find(&reviews).Error
	return reviews, err
}

func (repo reviewRepository) GetReviewByBookName(name string) ([]model.Review, error) {
	reviews := make([]model.Review, 0)
	err := repo.db.Table(model.Tables.Review).Where(model.ReviewTable.BookName+" = ?", name).Find(&reviews).Error
	return reviews, err
}

func (repo reviewRepository) GetReviewByID(id int) (model.Review, error) {
	var review model.Review
	err := repo.db.Table(model.Tables.Review).Where(model.ReviewTable.ID+" = ?", id).First(&review).Error
	return review, err
}
