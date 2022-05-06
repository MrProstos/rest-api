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

type Order struct {
	ID         uint
	Client_id  string `gorm:"primary_key;uniq"` //Принимает Client.Id
	Message_ID string `gorm:"primary_key:uniq"` //Принимает Message.Id
	Status     string
}

func (ord Order) Add(db *gorm.DB) error {
	return nil
}

type Message struct {
	ID    uint
	Id    string `gorm:"primary_key"` //Принимает Order.Message_ID
	Title string
	To    string
	Body  string
}
