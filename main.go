package main

import (
	"fmt"
	"log"
	"pustaka-api/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// Data Source Name (DSN)
	dsn := "egiwira:12345@tcp(127.0.0.1:3306)/pustaka_api?charset=utf8mb4&parseTime=True&loc=Local"

	// Open the database connection
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Database connection error: %v", err)
	} else {
		fmt.Println("Database connected")
	}

	// Initialize Gin router
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/rootHandler", handler.RootHandler)
		v1.GET("/helloHandler", handler.HelloHandler)
		v1.GET("/books/:id/title", handler.BooksHandler)
		v1.GET("/query", handler.QueryHandler)
		v1.POST("/books", handler.PostBooksHandler)
	}

	// Start the server
	if err := r.Run(); err != nil {
		log.Fatalf("Server run error: %v", err)
	}
}
