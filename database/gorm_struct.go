package database

import (
	"github.com/jinzhu/gorm"
)

type UnitNumber struct {
	gorm.Model
	Name       string `gorm:"column:name"`
	RomanValue string `gorm:"column:roman_value"`
}

type RomenNumerical struct {
	gorm.Model
	Key   string `gorm:"column:key"`
	Value int    `gorm:"column:value"`
}

type MaterialCredit struct {
	gorm.Model
	Name        string  `gorm:"column:name"`
	CreditValue float64 `gorm:"column:price"`
}
