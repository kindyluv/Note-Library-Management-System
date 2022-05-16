package data

type LibrarianRepository interface {
	SaveBook(book *Book) *Book
	UpdateBookByTitle(book *Book) *Book
}
