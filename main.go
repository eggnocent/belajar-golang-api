package main

import (
	"fmt"
	"log"
	"pustaka-api/book"
	"pustaka-api/handler"

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
		fmt.Println("Database connect")
	}
	if err := db.AutoMigrate(&book.Book{}); err != nil {
		log.Fatalf("Automigrate eror: %v", err)
	}

	//CREATE
	// book := book.Book{}
	// book.Title = "nama-nama pemain takraw ter unik"
	// book.Price = 120000
	// book.Discount = 11
	// book.Rating = 7
	// book.Description = "sports science"

	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat menambah data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }

	//READ DATA PERTAMA
	var bookfirst book.Book
	err = db.Debug().First(&bookfirst, 2).Error
	if err != nil {
		fmt.Println("====asdffghjkl=====")
		fmt.Println("tidak dapat mencari data")
		fmt.Println("====qwerttyuiop=====")
	}
	fmt.Println("title :", bookfirst.Title)
	fmt.Printf("book %v", bookfirst)

	//READ DATA TERAKHIR
	// var booklast book.Book

	// err = db.Debug().Last(&booklast).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// fmt.Println("title :", booklast.Title)
	// fmt.Printf("book %v", booklast)

	// READ DATA RANDOM
	// var bookrandomly book.Book
	// err = db.Debug().Take(&bookrandomly).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// fmt.Println("title :", bookrandomly.Title)
	// fmt.Printf("book %v", bookrandomly)

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
