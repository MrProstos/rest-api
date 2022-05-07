package db

import "github.com/jinzhu/gorm"

type Client struct {
	Id        uint
	Id_string string `gorm:"primary_key:uniq"`
	Firstname string
	Lastname  string
	Operator  string //Ввести Operator.Username
}

func (client *Client) Add(db *gorm.DB) error {
	db = db.AutoMigrate(&client)
	if db.Error != nil {
		return db.Error
	}
	db = db.Create(&client)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (client *Client) Change(db *gorm.DB) error {

	return nil
}

func (client *Client) Del(db *gorm.DB) error {
	db = db.Delete(&client, 1)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
