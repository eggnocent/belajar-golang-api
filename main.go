package main

import (
	"fmt"
	"net/http"

	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Stok     int    `json:"stok" binding:"required"`
	SubTitle string `json:"subtitle"`
}

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/rootHandler", handler.RootHandler)
	v1.GET("/helloHandler", handler.HelloHandler)
	v1.GET("/books/:id/title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", postBooksHandler)

	//v2 := r.Group("/v2")

	r.Run()
}

func postBooksHandler(c *gin.Context) {
	var bookInput BookInput

	if err := c.ShouldBindJSON(&bookInput); err != nil {
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

	c.JSON(http.StatusOK, gin.H{
		"title": bookInput.Title,
		"stok":  bookInput.Stok,
	})
}
