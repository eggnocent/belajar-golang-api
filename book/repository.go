package book

import (
	"fmt"

	"gorm.io/gorm"
)

// Repository interface defines the methods that any repository should implement.
type Repository interface {
	FindAll() ([]BookRequest, error)
	FindByID(ID int) (BookRequest, error)
	Create(book BookRequest) (BookRequest, error)
}

// repository struct implements the Repository interface.
type repository struct {
	db *gorm.DB
}

// NewRepository creates a new instance of the repository.
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// FindAll retrieves all BookRequest records from the database.
func (r *repository) FindAll() ([]BookRequest, error) {
	var books []BookRequest
	err := r.db.Find(&books).Error
	return books, err
}

// FindByID retrieves a single BookRequest record by its ID from the database.
func (r *repository) FindByID(ID int) (BookRequest, error) {
	var book BookRequest
	err := r.db.First(&book, ID).Error
	return book, err
}

func (r *repository) Create(book BookRequest) (BookRequest, error) {
	fmt.Printf("Creating book with details: %+v\n", book) // Log book details
	if err := r.db.Create(&book).Error; err != nil {
		fmt.Printf("Error creating book: %v\n", err) // Log error
		return BookRequest{}, err
	}
	fmt.Printf("Created Book with ID: %d\n", book.ID) // Log ID
	return book, nil
}
