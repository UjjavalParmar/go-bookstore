package models

import (
	"gorm.io/gorm"
	"github.com/ujjavalparmar/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	db.Create(&b) // No need for NewRecord in GORM v2
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	result := db.First(&getBook, Id) // Use First instead of Where().Find()
	return &getBook, result
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Delete(&book, ID) // GORM v2 Delete syntax
	return book
}

