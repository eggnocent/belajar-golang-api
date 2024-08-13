package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	if err := c.ShouldBindJSON(&bookinput); err != nil {
		c.JSON(http.StatusBadRequest, err)
		fmt.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title": bookinput.Title,
		"stok":  bookinput.Stok,
	})
}
