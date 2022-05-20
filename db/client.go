package db

import (
	"errors"
)

//Интерфейс для управления базой данных
type Db_manage interface {
	IsValid() error //Проверка на валидность данных
	Add() error     //Добавляет клиента в базу данных
	Update() error  //Обновление данных клиента
	Del() error     //Удаление клиента
}

//Структура данных Клиент содержащая поля такие же как и в базе данных
type Client struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Phone_num string `gorm:"unique;type:varchar;not null" json:"phone_num"`
	Firstname string `gorm:"type:varchar;not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar;not null" json:"lastname"`
	Birthday  string `gorm:"type:varchar;not null" json:"birthday"`
}

//Проверка на валидность данных
func (client *Client) IsValid() error {
	if len(client.Firstname) == 0 || len(client.Lastname) == 0 || len(client.Phone_num) == 0 {
		return errors.New("fields phone_num, firstname, lastname, birthday are required")
	}
	return nil
}

//Добавляет клиента в базуданных
func (client *Client) Add() error {
	if err := client.IsValid(); err != nil {
		return err
	}

	db := GetDB()
	if err := db.Create(&client); err.Error != nil {
		return err.Error
	}
	return nil
}

//Обновление данных клиентаdb/client.go
func (client *Client) Update() error {
	if err := client.IsValid(); err != nil {
		return err
	}

	db := GetDB()

	if err := db.Model(&client).Updates(&client); err.Error != nil {
		return err.Error
	}
	return nil
}

//Удаление клиента
func (client *Client) Del() error {
	db := GetDB()
	if err := db.Delete(&Client{}, client.ID); err.Error != nil {
		return err.Error
	}
	return nil
}
