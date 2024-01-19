package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Price       uint   `json:"price"`
}

func createBook(db *gorm.DB, book *Book) error {
	result := db.Create(&book)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func getBookById(db *gorm.DB, id int) *Book {
	var book Book

	result := db.Unscoped().First(&book, id)

	if result.Error != nil {
		log.Fatalf("Error when get book: %v", result.Error)
	}

	fmt.Println("Get book successful!")

	return &book
}

func getBooks(db *gorm.DB) []Book {
	var books []Book

	result := db.Unscoped().Find(&books)

	if result.Error != nil {
		log.Fatalf("Error when get book: %v", result.Error)
	}

	fmt.Println("Get book successful!")

	return books
}

func updateBook(db *gorm.DB, book *Book) error {
	result := db.Model(&book).Updates(book) // only update non zero fields, zero fields are ignored, refer https://gorm.io/docs/update.html#Update-Changed-Fields

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func deleteBook(db *gorm.DB, id uint) error {
	var book Book
	result := db.Delete(&book, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func searchBook(db *gorm.DB, bookName string) *Book {
	var book Book

	result := db.Where("name = ?", bookName).First(&book)
	if result.Error != nil {
		log.Fatalf("Search book failed: %v", result.Error)
	}

	return &book
}

func searchBooks(db *gorm.DB, bookName string) []Book {
	var books []Book

	result := db.Where("name = ?", bookName).Order("price asc").Find(&books)
	if result.Error != nil {
		log.Fatalf("Search book failed: %v", result.Error)
	}

	return books
}
