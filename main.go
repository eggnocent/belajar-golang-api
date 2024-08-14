package main

import (
	"fmt"
	"log"

	"pustaka-api/book"    // Adjust the import path according to your project structure
	"pustaka-api/handler" // Import your handler package

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "egiwira:12345@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	} else {
		fmt.Println("Database connected")
	}

	if err := db.AutoMigrate(&book.BookRequest{}); err != nil {
		log.Fatalf("Automigrate error: %v", err)
	}

	// Initialize repository, service, and handler
	bookRepository := book.NewRepository(db)
	bookService := book.NewService(bookRepository)
	bookHandler := handler.NewBookHandler(bookService)

	// Initialize Gin router
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/rootHandler", bookHandler.RootHandler)
		v1.GET("/helloHandler", bookHandler.HelloHandler)
		v1.GET("/books/:id/title", bookHandler.BooksHandler)
		v1.GET("/query", bookHandler.QueryHandler)
		v1.POST("/books", bookHandler.PostBooksHandler)
	}

	// Start the server
	if err := r.Run(); err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}
