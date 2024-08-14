package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookInput struct {
	Title    string `json:"title" binding:"required"`
	Stok     int    `json:"stok" binding:"required"`
	SubTitle string `json:"subtitle"`
}

func main() {
	r := gin.Default()

	r.GET("/rootHandler", rootHandler)
	r.GET("/helloHandler", helloHandler)
	r.GET("/books/:id/title", booksHandler)
	r.GET("/query", queryHandler)
	r.POST("/books", postBooksHandler)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":    bookInput.Title,
		"stok":     bookInput.Stok,
		"subtitle": bookInput.SubTitle,
	})
}
