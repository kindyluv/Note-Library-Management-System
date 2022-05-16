package utils

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
)

func DeleteCreatedModels(db *gorm.DB) func() {
	type entity struct {
		table   string
		KeyName string
		Key     interface{}
	}

	var entries []entity
	hookName := "cleanupHook"

	db.Callback().Create().After("gorm:create").Register(hookName,
		func(scope *gorm.Scope) {
			fmt.Printf("Inserted entries of %s with %s = %v\n",
				scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
			entries = append(entries, entity{table: scope.TableName(),
				KeyName: scope.PrimaryKey(), Key: scope.PrimaryKeyValue()})
		})
	return func() {
		defer db.Callback().Create().Remove(hookName)
		_, inTransaction := db.CommonDB().(*sql.Tx)
		tx := db
		if !inTransaction {
			tx = db.Begin()
		}
		for i := len(entries) - 1; i >= 0; i-- {
			entity := entries[i]
			fmt.Printf("Deleting entries from %s table with key %v\n", entity.table, tx.Table(entity.table).Where(entity.KeyName+"=?", entity.Key).Delete(""))
		}
		if !inTransaction {
			tx.Commit()
		}
	}
}
