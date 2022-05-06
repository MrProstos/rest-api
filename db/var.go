package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Operator struct {
	gorm.Model
	Username string
	Token    string
}

type DB_Manage interface {
	Add(*gorm.DB) error
	Change(*gorm.DB) error
	Del(*gorm.DB) error
}

func NewOperator(username string, token string) (op *Operator, err error) {
	if username == "" && token == "" {
		err = errors.New("missing data")
		return nil, err
	}
	op = &Operator{
		Username: username,
		Token:    token,
	}
	return
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
	gorm.Model
	Id        string
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
	gorm.Model
	Client_id  string //Принимает Client.Id
	Message_ID string //Принимает Message.Id
	Status     string
}

type Message struct {
	gorm.Model
	Id    string //Принимает Order.Message_ID
	Title string
	To    string
	Body  string
}
