package main

import (
	"fmt"
	"net/http"

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
	v1.GET("/roo`tHandler", rootHandler)
	v1.GET("/helloHandler", helloHandler)
	v1.GET("/books/:id/title", booksHandler)
	v1.GET("/query", queryHandler)
	v1.POST("/books", postBooksHandler)

	//v2 := r.Group("/v2")

	r.Run()
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"name": "Egi",
		"desc": "intern",
		"time": "10 months",
	})
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"divisi":   "Engineer",
		"position": "Backend",
		"materi":   "Belajar API pada golang",
	})
}

func booksHandler(c *gin.Context) {
	id := c.Param("id")
	title := c.Param("title")
	c.JSON(http.StatusOK, gin.H{
		"id":    id,
		"title": title,
	})
}

func queryHandler(c *gin.Context) {
	title := c.Query("title")
	stok := c.Query("stok")
	c.JSON(http.StatusOK, gin.H{
		"title": title,
		"stok":  stok,
	})
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
