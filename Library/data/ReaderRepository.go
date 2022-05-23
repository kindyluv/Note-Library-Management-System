package data

import (
	"github.com/jinzhu/gorm"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/dto"
	"log"
)

type ReaderRepository interface {
	CreateReader(reader dto.ReaderRequest) *dto.ReaderResponse
	FindReaderById(Id uint) *Reader
	FindReaderByUserName(name string) *Reader
	UpdateDetailsByUserName(name string) *Reader
	UpdateDetailsById(id uint) *Reader
	DeleteReaderById(id uint) string
	DeleteReaderByUserName(name string) string
}

type ReaderRepositoryImpl struct {
}

func (readerRepo ReaderRepositoryImpl) CreateReader(reader dto.ReaderRequest) *dto.ReaderResponse {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	foundReader := ModelMapper(reader)
	Db.Create(foundReader)
	Db.Where("ID = ? ", foundReader.ID).Find(&foundReader)
	log.Println("Created Reader is --> ", foundReader)
	response := new(dto.ReaderResponse)
	response.UserName = foundReader.UserName
	response.ID = foundReader.ID
	response.CreatedAt = foundReader.CreatedAt
	return response
}

func ModelMapper(reader dto.ReaderRequest) *Reader {
	bookReader := new(Reader)
	bookReader.FirstName = reader.FirstName
	bookReader.LastName = reader.LastName
	bookReader.Email = reader.Email
	bookReader.UserName = reader.UserName
	bookReader.Password = reader.Password
	return bookReader
}

func FindReaderById(id uint) *Reader {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	foundReader := &Reader{}
	Db.Where("ID = ? ", id).Find(foundReader)
	if foundReader != nil {
		return nil
	}
	return foundReader
}

func FindReaderByUserName(name string) *Reader {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	foundReader := &Reader{}
	Db.Where("UserName = ? ", name).Find(foundReader)
	if foundReader != nil {
		return nil
	}
	return foundReader
}

func (readerRepo *ReaderRepositoryImpl) FindReaderById(Id uint) *Reader {
	return FindReaderById(Id)
}

func (readerRepo *ReaderRepositoryImpl) FindReaderByUserName(name string) *Reader {
	return FindReaderByUserName(name)
}

func (readerRepo *ReaderRepositoryImpl) UpdateDetailsById(id uint) *Reader {
	foundReader := FindReaderById(id)
	log.Println("Reader to be updated is --> ", foundReader)
	var reader *Reader
	Db.First(&reader, foundReader)
	reader.ID = foundReader.ID
	reader.UserName = foundReader.UserName
	reader.ReaderAccount = foundReader.ReaderAccount
	reader.Password = foundReader.Password
	reader.UpdatedAt = foundReader.UpdatedAt
	Db.Update(reader)
	return reader
}

func (readerRepo ReaderRepositoryImpl) UpdateDetailsByUserName(name string) *Reader {
	foundReader := FindReaderByUserName(name)
	log.Println("Reader to be updated is --> ", foundReader)
	var reader *Reader
	Db.First(&reader, foundReader)
	reader.ID = foundReader.ID
	reader.UserName = foundReader.UserName
	reader.ReaderAccount = foundReader.ReaderAccount
	reader.Password = foundReader.Password
	reader.UpdatedAt = foundReader.UpdatedAt
	Db.Update(reader)
	return reader
}

func (readerRepo ReaderRepositoryImpl) DeleteReaderById(id uint) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	Db.Where("ID = ?", id).Delete(id)
	return "Reader successfully deleted"
}

func (readerRepo ReaderRepositoryImpl) DeleteReaderByUserName(name string) string {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	Db.Where("UserName = ?", name).Delete(name)
	return "Reader successfully deleted"
}
