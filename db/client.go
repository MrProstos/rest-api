package db

import (
	"errors"
)

type Db_manage interface {
	IsValid() error
	Add() error
	Update() error
	Del() error
}

type Client struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Phone_num string `gorm:"unique;type:varchar;not null" json:"phone_num"`
	Firstname string `gorm:"type:varchar;not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar;not null" json:"lastname"`
	Birthday  string `gorm:"type:varchar;not null" json:"birthday"`
}

func (client *Client) IsValid() error {
	if len(client.Firstname) == 0 || len(client.Lastname) == 0 || len(client.Phone_num) == 0 {
		return errors.New("fields phone_num, firstname, lastname, birthday are required")
	}
	return nil
}

func (client *Client) Add() error {
	if err := client.IsValid(); err != nil {
		return err
	}
	db := GetDB()
	err := db.Create(&client)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (client *Client) Update() error {
	if err := client.IsValid(); err != nil {
		return err
	}

	db := GetDB()
	if db.Error != nil {
		return db.Error
	}

	data := new(Client)
	data.ID = client.ID

	db.First(&data)
	data.Phone_num = client.Phone_num
	data.Firstname = client.Firstname
	data.Lastname = client.Lastname
	data.Birthday = client.Birthday
	db.Save(&data)

	return nil
}

func (client *Client) Del() error {
	db := GetDB()
	err := db.Delete(&client, 1)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
