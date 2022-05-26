package core

import (
	"errors"
	"github.com/MrProstos/rest-api/internal/db"
	"log"
)

// Client Структура данных Клиент содержащая поля такие же как и в базе данных
type Client struct {
	ID        uint   `gorm:"primaryKey"` // Первичный ключ autoincrement
	Phone_num string `gorm:"unique;type:varchar;not null" json:"phone_num"`
	Firstname string `gorm:"type:varchar;not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar;not null" json:"lastname"`
	Birthday  string `gorm:"type:varchar;not null" json:"birthday"` // Поменять на тип данных Date
	//OrderID   uint
	Orders []Order
}

func init() {
	getDB := db.GetDB()
	err := getDB.AutoMigrate(&Client{}, &Order{})
	if err != nil {
		log.Fatal(err)
	}
}

// IsValid Проверка на валидность данных
func (client *Client) IsValid() error {
	if len(client.Firstname) == 0 || len(client.Lastname) == 0 || len(client.Phone_num) == 0 {
		return errors.New("fields phone_num, firstname, lastname, birthday are required")
	}
	return nil
}

// Show SELECT * FROM clients WHERE phone_num = client.Phone_num
func (client *Client) Show() error {
	getDB := db.GetDB()
	if err := getDB.Where("phone_num = ?", client.Phone_num).First(&client); err != nil {
		return err.Error
	}
	return nil
}

// Add Добавляет клиента в базуданных
func (client *Client) Add() error {
	err := client.IsValid()
	if err != nil {
		return err
	}

	getDB := db.GetDB()
	if err := getDB.Create(&client); err.Error != nil {
		return err.Error
	}
	return nil
}

// Update Обновление данных клиента
func (client *Client) Update() error {
	err := client.IsValid()
	if err != nil {
		return err
	}

	getDB := db.GetDB()
	if err := getDB.Model(&client).Where("phone_num = ?", client.Phone_num).Updates(&client); err.Error != nil {
		return err.Error
	}
	return nil
}

// Del Удаление клиента
func (client *Client) Del() error {
	getDB := db.GetDB()
	err := getDB.Delete(&Client{}, client.Phone_num)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
