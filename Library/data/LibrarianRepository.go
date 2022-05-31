package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/dto"
	"log"
)

type LibrarianRepository interface {
	AddALibrarian(request *dto.LibrarianAccountRequest) *dto.LibrarianAccountResponse
	DeleteReaderAccountById(id uint) string
	DeleteReaderAccountByUserName(name string) string
}

type LibrarianRepositoryImpl struct {
}

func (LibrarianRepo LibrarianRepositoryImpl) AddALibrarian(request *dto.LibrarianAccountRequest) *dto.LibrarianAccountResponse {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	lib := ModelMap(request)
	Db.Create(ModelMap(request))
	Db.Where("ID = ? ", lib.ID).Find(&lib)
	log.Println("Created Librarian Account is --> ", lib)
	res := new(dto.LibrarianAccountResponse)
	res.UserName = lib.UserName
	res.ID = lib.ID
	res.CreatedAt = lib.CreatedAt
	return res
}

func ModelMap(request *dto.LibrarianAccountRequest) *Librarian {
	lib := new(Librarian)
	lib.FirstName = request.FirstName
	lib.LastName = request.LastName
	lib.UserName = request.UserName
	lib.Email = request.Email
	return lib
}

func (LibrarianRepo LibrarianRepositoryImpl) DeleteReaderAccountById(id uint) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	reader := new(ReaderRepositoryImpl)
	Db.Where("ID = ?", reader.FindReaderById(id)).Delete(id)
	return "Reader Account successfully deleted"
}

func (LibrarianRepo LibrarianRepositoryImpl) DeleteReaderAccountByUserName(name string) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	reader := new(ReaderRepositoryImpl)
	foundReader := reader.FindReaderByUserName(name)
	Db.Where("UserName = ?", foundReader).Delete(name)
	return "Reader Account successfully deleted"
}
