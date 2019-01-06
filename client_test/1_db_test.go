package client_test

import (
	"roman/database"
	"testing"

	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

type DatabaseSuiteTest struct {
	db *database.DB
}

var databaseSuiteTest DatabaseSuiteTest
var pathSqlitePath = "./roman_test.db"

func init() {
	databaseSuiteTest.db = database.NewDB()
}

func TestDatabaseSqlitePath(t *testing.T) {
	db := databaseSuiteTest.db
	db.SetPath(pathSqlitePath)
	assert.Equal(t, pathSqlitePath, db.GetPath())
}

func TestDatabaseGormConnection(t *testing.T) {
	db := databaseSuiteTest.db
	db.SetPath(pathSqlitePath)
	db.CreateConnection()

	gormDb := db.GetConnection()
	assert.NotEqual(t, gormDb, nil)
	assert.IsType(t, &gorm.DB{}, gormDb)
}
