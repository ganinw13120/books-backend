package service

import (
	"books-backend/model/response"
	"books-backend/repository"
)

type reviewService struct {
	bookRepository   repository.IBookRepository
	reviewRepository repository.IReviewRepository
}

type IReviewService interface {
	GetReviewList() (*response.GetReviewResponse, error)
}

func NewReviewController(
	bookRepository repository.IBookRepository,
	reviewRepository repository.IReviewRepository,
) reviewService {
	return reviewService{
		bookRepository:   bookRepository,
		reviewRepository: reviewRepository,
	}
}

func (service reviewService) GetReviewList() (*response.GetReviewResponse, error) {
	reviews, err := service.reviewRepository.GetReviewList()
	if err != nil {
		return nil, err
	}
	response := response.GetReviewResponse{
		Reviews: reviews,
	}
	return &response, nil
}
