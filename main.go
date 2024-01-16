package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydatabase"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Book{})
	fmt.Println("Migrate successful!")

	// Create
	// newBook := &Book{
	// 	Name:        "LL",
	// 	Author:      "Mikelopster",
	// 	Description: "LL",
	// 	Price:       600,
	// }
	// createBook(db, newBook)

	// GetBook
	// currentBook := getBookById(db, 1)
	// fmt.Println(currentBook)

	// // Update
	// currentBook := getBookById(db, 1)
	// currentBook.Name = "New Harry Potter"
	// currentBook.Price = 400

	// updateBook(db, currentBook)

	// Delete
	// deleteBook(db, 1)

	// search Book
	// currentBook := searchBook(db, "Harry Potter")
	// fmt.Println(currentBook)

	// search Book
	currentBooks := searchBooks(db, "LL")
	fmt.Println(currentBooks)

	for _, book := range currentBooks {
		fmt.Println(book.ID, book.Name, book.Author, book.Description, book.Price)
	}
}
