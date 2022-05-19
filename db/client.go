package db

import (
	"errors"
)

type Db_manage interface {
	Add() error
	Update() error
	Del() error
}

type Client struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Phone_num string `gorm:"unique;type:varchar" json:"phone_num"`
	Firstname string `gorm:"type:varchar" json:"firstname"`
	Lastname  string `gorm:"type:varchar" json:"lastname"`
	Birthday  string `gorm:"type:varchar" json:"birthday"`
}

func (client *Client) IsValid() error {
	if len(client.Firstname) == 0 || len(client.Lastname) == 0 || len(client.Phone_num) == 0 {
		return errors.New("fields Phone_num, Firstname, Lastname are required")
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

	data := new(Client)

	db := GetDB()
	err := db.First(&data, "id = ?", client.ID)
	if err.Error != nil {
		return err.Error
	}

	data.Phone_num = client.Phone_num
	data.Firstname = client.Firstname
	data.Lastname = client.Lastname
	data.Birthday = client.Birthday

	err = db.Save(&data)
	if err.Error != nil {
		return err.Error
	}
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
