package test

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/data"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/dto"
	"github.com/kindyluv/Note-Library-Management-System/tree/indev/Library/Library/utils"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

var (
	Db                               = data.Connect()
	readerRepo data.ReaderRepository = &data.ReaderRepositoryImpl{}
)

func TestThatReaderCanCreateNewReaderAccount(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	defer cleaner()
	reader := dto.ReaderRequest{FirstName: "Delight", LastName: "Lois", Email: "delight@gmail.com", UserName: "Adanne", Password: "ada___123"}
	savedReader := readerRepo.CreateReader(reader)
	log.Println("created at", savedReader.CreatedAt)
	assert.NotEmpty(t, reader.UserName, savedReader.UserName)
}

func TestThatReaderCanBeFoundById(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)

	defer cleaner()
	reader := dto.ReaderRequest{FirstName: "Delight", LastName: "Lois", Email: "delightada@gmail.com", UserName: "Adaora", Password: "ada___123"}
	savedReader := readerRepo.CreateReader(reader)
	log.Println("Reader Id --> ", savedReader.ID)
	assert.NotEmpty(t, savedReader)
	foundReader := readerRepo.FindReaderById(savedReader.ID)
	log.Println("found reader-->", foundReader)
	assert.Equal(t, foundReader.ID, savedReader.ID)
}

func TestThatReaderCanBeFoundByUserName(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)

	defer cleaner()
	reader := dto.ReaderRequest{FirstName: "Desire", LastName: "Shula", Email: "shul@gmail.com", UserName: "Ulma", Password: "bella"}
	savedReader := readerRepo.CreateReader(reader)
	assert.NotEmpty(t, savedReader)
	log.Println("Reader UserName --> ", savedReader.UserName)

	foundReader := readerRepo.FindReaderByUserName(savedReader.UserName)
	log.Println("Found Reader Username --> ", foundReader.UserName)
	log.Println("Saved Reader Username --> ", savedReader.UserName)
	assert.Equal(t, foundReader.UserName, savedReader.UserName)
}

func TestThatDetailsCanBeUpdatedByUserName(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	defer cleaner()

	reader := dto.ReaderRequest{FirstName: "Desire", LastName: "Shula", Email: "shula@gmail.com", UserName: "Uloma", Password: "bella"}
	savedReader := readerRepo.CreateReader(reader)
	assert.NotEmpty(t, savedReader)
	findReader := readerRepo.FindReaderById(savedReader.ID)
	assert.Equal(t, findReader.ID, savedReader.ID)
	foundReader := readerRepo.UpdateDetailsByUserName(findReader.UserName)
	log.Println("Found Reader ID --> ", foundReader.ID)
	log.Println("Find Reader ID --> ", findReader.ID)
	assert.Equal(t, foundReader.ID, findReader.ID)
}

func TestThatDetailsCanBeUpdatedById(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	defer cleaner()

	reader := dto.ReaderRequest{FirstName: "Desire", LastName: "Shula", Email: "shula@gmail.com", UserName: "Uloma", Password: "bella"}
	savedReader := readerRepo.CreateReader(reader)
	assert.NotEmpty(t, savedReader)

	findReader := readerRepo.FindReaderById(savedReader.ID)
	log.Println("found reader-->", findReader)
	assert.Equal(t, findReader.ID, savedReader.ID)

	foundReader := readerRepo.UpdateDetailsById(savedReader.ID)
	assert.Equal(t, foundReader.ID, findReader.ID)
}

func TestThatReaderCanBeDeletedById(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	defer cleaner()
	reader := dto.ReaderRequest{FirstName: "Diva", LastName: "Shula", Email: "o-i@gmail.com", UserName: "k-k", Password: "bella"}
	savedReader := readerRepo.CreateReader(reader)
	assert.NotEmpty(t, savedReader)
	log.Println("Find all --> ", len(readerRepo.FindAllReaders()))
	assert.Equal(t, 3, len(readerRepo.FindAllReaders()))
	readerRepo.DeleteReaderById(savedReader.ID)
	assert.Equal(t, 2, len(readerRepo.FindAllReaders()))
}
