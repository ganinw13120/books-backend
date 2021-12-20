package service

import (
	"books-backend/model/response"
	"books-backend/repository"
)

type reviewService struct {
	reviewRepository repository.IReviewRepository
}

type IReviewService interface {
	GetAllReviewList() (*response.GetReviewListResponse, error)
	GetReviewByBookName(name string) (*response.GetReviewListResponse, error)
	GetReviewByID(id int) (*response.GetReviewResponse, error)
}

func NewReviewController(
	reviewRepository repository.IReviewRepository,
) reviewService {
	return reviewService{
		reviewRepository: reviewRepository,
	}
}

func (service reviewService) GetReviewByBookName(name string) (*response.GetReviewListResponse, error) {
	reviews, err := service.reviewRepository.GetReviewByBookName(name)
	if err != nil {
		return nil, err
	}
	response := response.GetReviewListResponse{
		Reviews: reviews,
	}
	return &response, nil
}

func (service reviewService) GetReviewByID(id int) (*response.GetReviewResponse, error) {
	reviews, err := service.reviewRepository.GetReviewByID(id)
	if err != nil {
		return nil, err
	}
	response := response.GetReviewResponse{
		Review: reviews,
	}
	return &response, nil
}

func (service reviewService) GetAllReviewList() (*response.GetReviewListResponse, error) {
	reviews, err := service.reviewRepository.GetAllReviewList()
	if err != nil {
		return nil, err
	}
	response := response.GetReviewListResponse{
		Reviews: reviews,
	}
	return &response, nil
}
