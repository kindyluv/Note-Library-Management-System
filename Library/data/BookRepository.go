package data

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/dto"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/utils"
	"log"
)

type BookRepository interface {
	SaveBook(book dto.BookRequest) *dto.BookResponse
	FindByBookId(bookId uint) *Book
	FindByBookTitle(bookTitle string) *Book
	FindByBookIsbn(bookIsbn string) *Book
	FindByBookAuthor(bookAuthor string) *Book
	FindByBookEdition(bookEdition string) *Book
	FindAllBooks() []*Book
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
	config, err = utils.LoadConfig("/home/precious/Desktop/dev/LibrarySystem/LibrarySystem/Library/utils")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("config-->", config)
	Db, err = gorm.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error connecting to db-->", err)
	}
	log.Println("db-->", Db)
	Db.AutoMigrate(&Book{}, &Reader{}, &Account{}, &Librarian{}, Author{})
	log.Println("connected " + "to db")
	return Db
}

type BookRepositoryImpl struct {
}

func (bookRepo *BookRepositoryImpl) SaveBook(request dto.BookRequest) *dto.BookResponse {
	Db := Connect()
	defer func(Db *gorm.DB) {
		err := Db.Close()
		if err != nil {

		}
	}(Db)
	book := fileUploadRequest(request)
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
	Db.Where("ID = ? ", bookId).Find(&savedBooks)

	if savedBooks == nil {
		return nil
	}
	return savedBooks
}

func (bookRepo *BookRepositoryImpl) FindByBookId(bookId uint) *Book {
	return FindByBookId(bookId)
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
	Db.Where("Title = ? ", bookTitle).Find(savedBooks)
	if savedBooks == nil {
		return nil
	}
	return savedBooks
}

func (bookRepo *BookRepositoryImpl) FindByBookTitle(bookTitle string) *Book {
	return FindByBookTitle(bookTitle)
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
	if savedBooks == nil {
		return nil
	}
	return savedBooks
}

func (bookRepo *BookRepositoryImpl) FindByBookIsbn(bookIsbn string) *Book {
	return FindByBookIsbn(bookIsbn)
}

func (bookRepo *BookRepositoryImpl) FindByBookAuthor(bookAuthor string) *Book {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	savedBook := &Book{}
	Db.Where("Author = ? ", bookAuthor).First(savedBook)
	if savedBook == nil {
		return nil
	}
	return savedBook
}

func (bookRepo *BookRepositoryImpl) FindByBookEdition(bookEdition string) *Book {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	savedBook := &Book{}
	Db.Where("Edition = ? ", bookEdition).First(savedBook)
	if savedBook == nil {
		return nil
	}
	return savedBook
}

func (bookRepo *BookRepositoryImpl) FindAllBooks() []*Book {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	var books []*Book
	Db.Find(books)
	return books
}
func (bookRepo *BookRepositoryImpl) UpdateByBookId(bookId uint) string {
	foundBook := FindByBookId(bookId)
	log.Println("Book to be updated --> ", foundBook)
	var book Book
	Db.First(&book, foundBook)
	book.ID = foundBook.ID
	book.Title = foundBook.Title
	book.ISBN = foundBook.ISBN
	book.Edition = foundBook.Edition
	book.BookUrl = foundBook.BookUrl
	book.Author = foundBook.Author
	book.UpdatedAt = foundBook.UpdatedAt
	Db.Update(book)
	return "Book with this title " + foundBook.Edition + " has been updated"
}

func (bookRepo *BookRepositoryImpl) UpdateBookByTitle(title string) string {
	foundBook := FindByBookTitle(title)
	log.Println("Book to be updated --> ", foundBook)
	var book Book
	Db.First(&book, foundBook)
	book.ID = foundBook.ID
	book.Title = foundBook.Title
	book.ISBN = foundBook.ISBN
	book.Edition = foundBook.Edition
	book.BookUrl = foundBook.BookUrl
	book.Author = foundBook.Author
	book.UpdatedAt = foundBook.UpdatedAt
	Db.Update(book)
	return "Book with this title " + foundBook.Title + " has been updated"
}

func (bookRepo *BookRepositoryImpl) UpdateBookByIsbn(bookIsbn string) string {
	foundBook := FindByBookIsbn(bookIsbn)
	log.Println("Book to be updated --> ", foundBook)
	var book Book
	Db.First(&book, foundBook)
	book.ID = foundBook.ID
	book.Title = foundBook.Title
	book.ISBN = foundBook.ISBN
	book.Edition = foundBook.Edition
	book.BookUrl = foundBook.BookUrl
	book.Author = foundBook.Author
	book.UpdatedAt = foundBook.UpdatedAt
	Db.Update(book)
	return "Book with this title " + foundBook.ISBN + " has been updated"
}

func (bookRepo *BookRepositoryImpl) DeleteByBookId(bookId uint) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	Db.Where("ID = ? ", bookId).Delete(&Book{})
	return "Book successfully deleted"
}

func (bookRepo *BookRepositoryImpl) DeleteBookByTitle(title string) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	Db.Where("Title = ? ", title).Delete(&Book{})
	return "Book successfully deleted"
}

func (bookRepo *BookRepositoryImpl) DeleteBookByIsbn(bookIsbn string) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	Db.Where("ISBN = ? ", bookIsbn).Delete(&Book{})
	return "Book successfully deleted"
}
