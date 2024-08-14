package handler

import (
	"fmt"
	"net/http"

	"pustaka-api/book"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type booksHandler struct {
	bookService book.Service
}

func NewBookHandler(bookService book.Service) *booksHandler {
	return &booksHandler{bookService}
}

func (h *booksHandler) RootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Egi",
		"desc": "intern",
		"time": "10 months",
	})
}

func (h *booksHandler) HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"divisi":   "Engineer",
		"position": "Backend",
		"materi":   "Belajar API pada golang",
	})
}

func (h *booksHandler) BooksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func (h *booksHandler) QueryHandler(c *gin.Context) {
	title := c.Query("title")
	stok := c.Query("stok")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"stok":  stok,
	})
}

func (h *booksHandler) PostBooksHandler(c *gin.Context) {
	var bookRequest book.BookRequest

	if err := c.ShouldBindJSON(&bookRequest); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			var errors []string
			for _, e := range validationErrors {
				errors = append(errors, fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag()))
			}
			c.JSON(http.StatusBadRequest, gin.H{"errors": errors})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book, err := h.bookService.Create(bookRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"errors": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": book,
	})
}
