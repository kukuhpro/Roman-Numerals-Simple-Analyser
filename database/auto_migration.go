package database

import "log"

func AutoMigration() {
	log.Println("Running Auto Migration on sqlite...")
	model := NewModel()
	db := model.Db.GetConnection()
	db.AutoMigrate(&UnitNumber{}, &RomenNumerical{}, &MaterialCredit{})
}
