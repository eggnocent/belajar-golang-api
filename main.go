package main

import (
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1")
	v1.GET("/rootHandler", handler.RootHandler)
	v1.GET("/helloHandler", handler.HelloHandler)
	v1.GET("/books/:id/title", handler.BooksHandler)
	v1.GET("/query", handler.QueryHandler)
	v1.POST("/books", handler.PostBooksHandler)

	//v2 := r.Group("/v2")

	r.Run()
}
