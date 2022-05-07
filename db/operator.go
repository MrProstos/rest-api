package db

import (
	"github.com/jinzhu/gorm"
)

type Operator struct {
	ID       uint `gorm:"primary_key:uniq"`
	Username string
	Token    string
}

func (oper *Operator) Add(db *gorm.DB) error {
	db = db.AutoMigrate(&oper)
	if db.Error != nil {
		return db.Error
	}
	db = db.Create(&oper)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (oper *Operator) Del(db *gorm.DB) error {
	db = db.Delete(&oper, 1)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
