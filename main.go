package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookInput struct {
	Title string `json:"title" binding:"required"`
	Stok  int    `json:"stok" binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/books", postBooksHandler)

	r.Run() // Jalankan server di localhost:8080
}

func postBooksHandler(c *gin.Context) {
	var bookinput BookInput

	err := c.ShouldBindJSON(&bookinput)
	if err != nil {
		// Cek apakah error adalah error validasi
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validationErrors {
				errorMessage := fmt.Sprintf("error on field %s, condition: %s", e.Field(), e.ActualTag())
				c.JSON(http.StatusBadRequest, errorMessage)
				return
			}

		}

		// Kembalikan pesan kesalahan umum
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookinput.Title,
		"stok":  bookinput.Stok,
	})
}
