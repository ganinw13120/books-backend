package service

import (
	"books-backend/model"
	"books-backend/repository"
)

type bookService struct {
	bookRepository   repository.IBookRepository
	reviewRepository repository.IReviewRepository
}

type IBookService interface {
	GetBooks(request model.GetBooksRequest) (*model.GetBooksResponse, error)
}

func NewBookController(
	bookRepository repository.IBookRepository,
	reviewRepository repository.IReviewRepository,
) bookService {
	return bookService{bookRepository: bookRepository}
}

// Get Books data
func (service bookService) GetBooks(request model.GetBooksRequest) (*model.GetBooksResponse, error) {
	return nil, nil
	// books := make([]model.Book, 0)

	// err := service.bookRepository.GetBooksData(request.Name, &books)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return nil, err
	// }
	// response := model.GetBooksResponse{
	// 	Data: books,
	// }
	// return &response, nil
}
