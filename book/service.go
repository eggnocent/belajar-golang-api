package book

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
	// Explicitly cast each field even though they seem to match in type.
	// This can sometimes help in cases where Go's type system expects explicit confirmation.
	book := BookRequest{
		Title:       string(bookInput.Title),       // Cast string to string
		Description: string(bookInput.Description), // Cast string to string
		Price:       int(bookInput.Price),          // Cast int to int
		Discount:    int(bookInput.Discount),       // Cast int to int
		Rating:      int(bookInput.Rating),         // Cast int to int
	}

	// Save the new book to the repository
	newBook, err := s.repository.Create(book)
	if err != nil {
		return BookRequest{}, err // Return an empty BookRequest and the error if creation fails
	}

	return newBook, nil // Return the newly created book and nil error
}
