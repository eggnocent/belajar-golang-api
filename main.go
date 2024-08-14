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

	bookRepository := book.NewRepository(db)

	book := book.Book{
		Title:       "Orkes pensil alis",
		Description: "The best gigs performing",
		Price:       100000,
		Rating:      5,
		Discount:    0,
	}

	bookRepository.Create(book)

	// FIND ALL
	// books, err := bookRepository.FindAll()

	// FIND BY ID
	//books, err := bookRepository.FindByID(2)
	//fmt.Println("Title:", books.Title)

	//CREATE
	// book := book.Book{}
	// book.Title = "jenis hewan bermata kaki"
	// book.Price = 1312
	// book.Discount = 1232
	// book.Rating = 5
	// book.Description = "animal say first"
	// err = db.Create(&book).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat menambah data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }

	//READ DATA PERTAMA
	// var bookfirst book.Book
	// err = db.Debug().First(&bookfirst).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// fmt.Println("title :", bookfirst.Title)
	// fmt.Printf("book %v", bookfirst)

	// READ DATA PRIMARY KEY
	// var bookfirst book.Book
	// err = db.Debug().First(&bookfirst, 1).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// fmt.Println("title :", bookfirst.Title)
	// fmt.Printf("book %v", bookfirst)

	// READ DATA ""
	// var bookfirst book.Book
	// err = db.Debug().First(&bookfirst, "1").Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// fmt.Println("title :", bookfirst.Title)
	// fmt.Printf("book %v", bookfirst)

	// READ DATA DENGAN ID YANG DITENTUKAN
	// var books []book.Book // Declare a slice to hold multiple Book records
	// err = db.Debug().Find(&books, []int{1, 2}).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// } else {
	// 	for _, book := range books {
	// 		fmt.Println("Title:", book.Title)
	// 		fmt.Printf("Book: %v\n", book)
	// 	}
	// }

	// MENCARI BERDASARKN STRING
	// var books []book.Book
	// err = db.Debug().Where("title = ?", "bagaimana cara magicom tau jika nasi sudah matang").Find(&books).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// } else {
	// 	for _, book := range books {
	// 		fmt.Println("Title:", book.Title)
	// 		fmt.Printf("Book: %v\n", book)
	// 	}
	// }

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
	//   var bookrandomly book.Book

	//   err = db.Debug().Take(&bookrandomly).Error
	//   if err != nil {
	//   	fmt.Println("====asdffghjkl=====")
	//   	fmt.Println("tidak dapat mencari data")
	//   	fmt.Println("====qwerttyuiop=====")
	//   }
	//   fmt.Println("title :", bookrandomly.Title)
	//   fmt.Printf("book %v", bookrandomly)

	// UPDATE DATA
	// var bookupdate book.Book

	// err = db.Debug().Where("id = ?", 2).First(&bookupdate).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// bookupdate.Title = "Daftar pemain bola kasti unik"
	// err = db.Save(&bookupdate).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat meng update data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }

	// DELETE DATA
	// var bookdelete book.Book

	// err = db.Debug().Where("id = ?", 3).First(&bookdelete).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat mencari data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
	// err = db.Delete(&bookdelete).Error
	// if err != nil {
	// 	fmt.Println("====asdffghjkl=====")
	// 	fmt.Println("tidak dapat menghapus data")
	// 	fmt.Println("====qwerttyuiop=====")
	// }
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
