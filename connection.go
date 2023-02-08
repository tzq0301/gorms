package gorms

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// InMemorySqlite create in memory sqlite database
//
// InMemorySqlite can be used for mocking database while unit testing
func InMemorySqlite() (*gorm.DB, error) {
	return gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
}

// InMemorySqliteWithInitSQL create in memory sqlite database with init SQL
func InMemorySqliteWithInitSQL(sql string, values ...interface{}) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.Exec(sql, values...).Error
	if err != nil {
		return nil, err
	}

	return db, nil
}
