package database

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var instancesDB *gorm.DB
var once sync.Once

type DB struct {
	gormDb     *gorm.DB
	sqlitePath string
}

func (d *DB) SetPath(sqlitePath string) {
	d.sqlitePath = sqlitePath
}

func (d *DB) GetPath() string {
	return d.sqlitePath
}

func (d *DB) CreateConnection() {
	once.Do(func() {
		db, err := gorm.Open("sqlite3", d.sqlitePath)
		if err != nil {
			log.Fatalln(err)
		}
		instancesDB = db
	})
	d.gormDb = instancesDB
}

func (d *DB) GetConnection() *gorm.DB {
	return d.gormDb
}

func NewDB() *DB {
	var db DB
	return &db
}
