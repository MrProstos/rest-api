package db

import (
	"github.com/jinzhu/gorm"
)

type Operator struct {
	ID       uint
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

func (oper *Operator) Update(data *Operator, db *gorm.DB) error {
	err := db.Model(&oper).Update(data)
	if err.Error != nil {
		return err.Error
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
