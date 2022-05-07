package db

import "github.com/jinzhu/gorm"

type Order struct {
	ID         uint
	Client_id  string `gorm:"primary_key;uniq"` //Принимает Client.Id
	Message_ID string `gorm:"primary_key:uniq"` //Принимает Message.Id
	Status     string
}

func (ord Order) Add(db *gorm.DB) error {
	return nil
}
