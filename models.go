package main

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func createBook(db *gorm.DB, book *Book) {
	result := db.Create(&book)

	if result.Error != nil {
		log.Fatalf("Error when create book: %v", result.Error)
	}

	fmt.Println("Create book successful!")
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

func updateBook(db *gorm.DB, book *Book) {
	result := db.Save(&book)

	if result.Error != nil {
		log.Fatalf("Error when update book: %v", result.Error)
	}

	fmt.Println("Update book successful!")
}

func deleteBook(db *gorm.DB, id uint) {
	var book Book
	result := db.Delete(&book, id)

	if result.Error != nil {
		log.Fatalf("Error when delete book: %v", result.Error)
	}

	fmt.Println("Delete book successful!")
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
