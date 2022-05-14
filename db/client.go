package db

import (
	"github.com/jinzhu/gorm"
)

type Client struct {
	Phone_num string
	Firstname string
	Lastname  string
	Birthday  string
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

func (client *Client) Update(data *Client, db *gorm.DB) error {
	err := db.Model(&client).Update(data)
	if err.Error != nil {
		return err.Error
	}
	return nil

}

func (client *Client) Del(db *gorm.DB) error {
	db = db.Delete(&client, 1)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
