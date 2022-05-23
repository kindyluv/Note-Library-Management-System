package test

import (
	"github.com/jinzhu/gorm"
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

func setUp() []*dto.ReaderRequest {
	return []*dto.ReaderRequest{
		{
			FirstName: "Lois",
			LastName:  "Precious",
			UserName:  "LoisB",
			Password:  "1234lois",
		},
		{
			FirstName: "Amara",
			LastName:  "Delight",
			UserName:  "Adaora",
			Password:  "123___",
		},
	}
}

func TestThatReaderCanCreateNewReaderAccount(t *testing.T) {
	cleaner := utils.DeleteCreatedModels(Db)
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(Db)
	defer cleaner()
	reader := dto.ReaderRequest{FirstName: "Delight", LastName: "Lois", Email: "delight@gmail.com", UserName: "Adaora", Password: "ada___123"}
	savedReader := readerRepo.CreateReader(reader)
	assert.NotEmpty(t, reader.UserName, savedReader.UserName)
}

//
//func TestThatReaderCanBeFoundById() {
//
//}
