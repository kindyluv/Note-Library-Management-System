package data

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/dto"
	"log"
)

type ReaderRepository interface {
	CreateReader(reader dto.ReaderRequest) *Reader
	FindReaderById(Id uint) *Reader
	FindReaderByUserName(name string) *Reader
	FindAllReaders() []*Reader
	UpdateDetailsByUserName(name string) *Reader
	UpdateDetailsById(id uint) *Reader
	DeleteReaderById(id uint) string
	DeleteReaderByUserName(name string) string
}

type ReaderRepositoryImpl struct {
}

func (readerRepo ReaderRepositoryImpl) CreateReader(reader dto.ReaderRequest) *Reader {
	Db := Connect()
	//defer func(db *gorm.DB) {
	//	err := db.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}(Db)
	foundReader := ModelMapper(reader)
	Db.Create(foundReader)
	Db.Where("ID = ? ", foundReader.ID).Find(&foundReader)
	log.Println("Created Reader is --> ", foundReader)
	return foundReader
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
	Db.Where("ID=? ", id).Find(&foundReader)
	log.Println("ID ............ ", foundReader.ID)
	if foundReader == nil {
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
	Db.Where("UserName=? ", name).Find(&foundReader)
	if foundReader == nil {
		return nil
	}
	return foundReader
}

func (readerRepo *ReaderRepositoryImpl) FindReaderById(Id uint) *Reader {
	return FindReaderById(Id)
}

func (readerRepo *ReaderRepositoryImpl) FindReaderByUserName(name string) *Reader {
	Db := Connect()
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	reader := Reader{}
	Db.First(&reader, "user_name=?", name)
	log.Println("here--", reader)
	return &reader
}

func (readerRepo *ReaderRepositoryImpl) FindAllReaders() []*Reader {
	Db := Connect()
	defer func(DB *gorm.DB) {
		err := DB.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	var readers []*Reader
	Db.Find(&readers)
	return readers
}

func (readerRepo *ReaderRepositoryImpl) UpdateDetailsById(id uint) *Reader {
	foundReader := &Reader{}
	Db.Where("id = ? ", id).Find(&foundReader)
	log.Println("Reader to be updated is --> ", foundReader)
	var reader *Reader
	Db.Model(&reader).Where("id = ?", id).Updates(Reader{FirstName: foundReader.FirstName, LastName: foundReader.LastName, Email: foundReader.Email, UserName: foundReader.UserName, Password: foundReader.Password})
	Db.Save(reader)
	return reader
}

func (readerRepo ReaderRepositoryImpl) UpdateDetailsByUserName(name string) *Reader {
	foundReader := FindReaderByUserName(name)
	log.Println("Reader to be updated is --> ", foundReader)
	var reader *Reader
	Db.First(&reader, foundReader)
	//Db.Model(Rea)
	Db.Model(&reader).Updates(Reader{FirstName: foundReader.FirstName, LastName: foundReader.LastName, Email: foundReader.Email, UserName: foundReader.UserName, Password: foundReader.Password})
	return reader
}

func (readerRepo ReaderRepositoryImpl) DeleteReaderById(id uint) string {
	Db := Connect()
	//Db.Close().Error()
	Db.Where("id = ?", id).Delete(&Reader{})
	return "Reader successfully deleted"
}

func (readerRepo ReaderRepositoryImpl) DeleteReaderByUserName(name string) string {
	Db := Connect()
	defer Db.Close().Error()
	Db.Where("UserName = ?", name).Delete(&Reader{})
	return "Reader successfully deleted"
}
