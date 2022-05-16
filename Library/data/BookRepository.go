package data

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/dto"
	"log"
)
import "github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/utils"

type BookRepository interface {
	SaveBook(book dto.BookRequest) *dto.BookResponse
	FindByBookId(bookId uint) *Book
	FindByBookTitle(bookTitle string) *Book
	FindByBookIsbn(bookIsbn string) *Book
	UpdateByBookId(bookId uint) string
	UpdateBookByTitle(title string) string
	UpdateBookByIsbn(bookIsbn string) string
	DeleteByBookId(bookId uint) string
	DeleteBookByTitle(title string) string
	DeleteBookByIsbn(bookIsbn string) string
}

var (
	config utils.Config
	Db     *gorm.DB
	err    error
)

func Connect() *gorm.DB {
	config, err = utils.LoadConfig("/home/precious/Desktop/dev/LibrarySystem/LibrarySystem")
	if err != nil {
		log.Fatal(err)
	}
	Db, err = gorm.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	Db.AutoMigrate(&Book{}, &Reader{}, &Account{}, &Librarian{}, Author{})
	log.Println("connected " + "to db")
	return Db
}

type BookRepositoryImpl struct {
}

func (bookRepo *BookRepositoryImpl) SaveBook(book *Book) *dto.BookResponse {
	Db := Connect()
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(Db)
	savedBook := &Book{}
	Db.Create(book)
	Db.Where("ID = ?", book.ID).Find(&savedBook)
	log.Println("Saved book is --> ", savedBook)
	response := new(dto.BookResponse)
	response.Url = savedBook.BookUrl
	response.Title = savedBook.Title
	return response
}

func fileUploadRequest(request dto.BookRequest) *Book {
	book := new(Book)
	book.Title = request.Title
	book.BookUrl = request.FileUrl
	book.ISBN = request.ISBN
	book.Edition = request.Edition
	s := request.Author
	err = json.Unmarshal([]byte(s), &book.Author)
	if err != nil {
		return nil
	}
	return book
}

func FindByBookId(bookId uint) *Book {
	Db := Connect()
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	savedBooks := &Book{}
	Db.Where("Book ID = ? ", bookId).Find(&savedBooks)

	if savedBooks == nil {
		return nil
	}
	return savedBooks
}

func FindByBookTitle(bookTitle string) *Book {
	Db := Connect()
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	savedBooks := &Book{}
	Db.Where("Book Title = ? ", bookTitle).Find(savedBooks)
	if savedBooks != nil {
		return nil
	}
	return savedBooks
}
func FindByBookIsbn(bookIsbn string) *Book {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	savedBooks := &Book{}
	Db.Where("ISBN = ? ", bookIsbn).Find(savedBooks)
}
