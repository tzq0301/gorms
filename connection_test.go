package gorms

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInMemorySqlite(t *testing.T) {
	a := assert.New(t)

	db, err := InMemorySqlite()
	a.NoError(err)

	a.NoError(db.Exec(`
CREATE TABLE user (id int);

INSERT INTO user (id) VALUES (1), (2), (3);
`).Error)

	var count int64
	db.Table("user").Count(&count)
	a.EqualValues(3, count)
}

func TestInMemorySqliteWithInitSQL(t *testing.T) {
	a := assert.New(t)

	db, err := InMemorySqliteWithInitSQL(`
CREATE TABLE user (id int);

INSERT INTO user (id) VALUES (1), (2), (3);
`)
	a.NoError(err)

	var count int64
	db.Table("user").Count(&count)
	a.EqualValues(3, count)

	_, err = InMemorySqliteWithInitSQL(`
CREATE TABLE user (id int);
CREATE TABLE user (id int, name VARCHAR(20));
`)
	a.Error(err)
}
