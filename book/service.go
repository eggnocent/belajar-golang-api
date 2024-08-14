package book

import "fmt"

type Service interface {
	FindAll() ([]BookRequest, error)
	FindByID(ID int) (BookRequest, error)
	Create(bookInput BookInput) (BookRequest, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]BookRequest, error) {
	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (BookRequest, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookInput BookInput) (BookRequest, error) {
	book := BookRequest{
		Title:       bookInput.Title,
		Description: bookInput.Description,
		Price:       bookInput.Price,
		Discount:    bookInput.Discount,
		Rating:      bookInput.Rating,
	}

	fmt.Printf("Creating book with details: %+v\n", book) // Logging book details

	newBook, err := s.repository.Create(book)
	if err != nil {
		return BookRequest{}, err
	}

	return newBook, nil
}
