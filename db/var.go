package db

import (
	"errors"

	"github.com/jinzhu/gorm"
)

type Operator struct {
	Username string
	Token    string
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

func (oper *Operator) Add(db *gorm.DB) string {
	db.AutoMigrate(oper)
	db.Create(oper)
	return "Success"
}

func (oper *Operator) Del(db *gorm.DB) string {
	db.Delete(oper)
	return "Success"
}

type Client struct {
	Id        string
	Firstname string
	Lastname  string
	Operator  string //Ввести Operator.Username
}

type Order struct {
	Client_id  string //Принимает Client.Id
	Message_ID string //Принимает Message.Id
	Status     string
}

type Message struct {
	Id    string //Принимает Order.Message_ID
	Title string
	To    string
	Body  string
}
