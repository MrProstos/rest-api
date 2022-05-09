package db

import "github.com/jinzhu/gorm"

type Order struct {
	ID        uint
	Client_id string `gorm:"foreignKey"`
	Title     string
	To        string
	Body      string
	Status    string
}

func (ord Order) Add(db *gorm.DB) error {
	db = db.AutoMigrate(&ord)
	if db.Error != nil {
		return db.Error
	}
	db = db.Create(&ord)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (ord Order) Update(data *Operator, db *gorm.DB) error {
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
