package db

import "github.com/jinzhu/gorm"

type Order struct {
	Client_id uint `gorm:"not null"`
	Title     string
	To        string
	Body      string
	Status    uint
}

func (ord Order) Add(db *gorm.DB) error {
	db = db.Create(&ord)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (ord Order) Update(data *Order, db *gorm.DB) error {
	err := db.Model(&ord).Update(data)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Del(db *gorm.DB) error {
	db = db.Delete(&ord, 1)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
