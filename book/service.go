package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(book BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	books, err := s.repository.FindAll()
	return books, err

	// return s.repository.FindAll() VERSI SIMPLE
}

func (s *service) FindByID(ID int) (Book, error) {
	book, err := s.repository.FindByID(ID)
	return book, err
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	// Assuming Stok is an int or can be directly converted to int
	book := Book{
		Title: bookRequest.Title,
		Price: bookRequest.Stok, // Directly using Stok if it's already an int
	}

	newBook, err := s.repository.Create(book)
	if err != nil {
		return Book{}, err // Return an empty Book and the error if creation fails
	}

	return newBook, nil // Return the newly created book and nil error
}
