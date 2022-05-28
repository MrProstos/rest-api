package db

import (
	"errors"
)

type ManageDb interface {
	IsValid() error //Проверка на валидность данных
	Show() error    //Получить данные клиента по номеру телефона
	Add() error     //Добавляет клиента в базу данных
	Update() error  //Обновление данных клиента
	Del() error     //Удаление клиента
}

// Client Структура данных Клиент содержащая поля такие же как и в базе данных
type Client struct {
	//ID        uint    `gorm:"primaryKey" json:"id"` // Первичный ключ autoincrement . ПОМЕНЯТЬ НА НОМЕР ТЕЛЕФОНА
	PhoneNum  string  `gorm:"primaryKey;unique;type:varchar;not null" json:"phone_num"`
	Firstname string  `gorm:"type:varchar;not null" json:"firstname"`
	Lastname  string  `gorm:"type:varchar;not null" json:"lastname"`
	Birthday  string  `gorm:"type:date;not null" json:"birthday"` // Поменять на тип данных Date
	Orders    []Order `gorm:"foreignKey:PhoneNum"`
}

type Order struct {
	ID       uint   `gorm:"primaryKey"`
	PhoneNum string `gorm:"type:varchar;not null"`
	To       string
	Body     string
	Status   uint `gorm:"not null;default:1"`
}

// IsValid Проверка на валидность данных
func (client *Client) IsValid() error {
	if len(client.Firstname) == 0 || len(client.Lastname) == 0 || len(client.PhoneNum) == 0 {
		return errors.New("fields phone_num, firstname, lastname, birthday are required")
	}
	return nil
}

// Show Вывод клиентов по номеру телефона
func (client *Client) Show() error {
	getDB := GetDB()
	if err := getDB.Where("phone_num = ?", client.PhoneNum).First(&client); err != nil {
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

	getDB := GetDB()
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

	getDB := GetDB()
	if err := getDB.Model(&client).Where("phone_num = ?", client.PhoneNum).Updates(&client); err.Error != nil {
		return err.Error
	}
	return nil
}

// Del Удаление клиента
func (client *Client) Del() error {
	getDB := GetDB()

	err := getDB.Exec("DELETE FROM public.orders WHERE orders.phone_num = ?", client.PhoneNum)
	if err.Error != nil {
		return err.Error
	}

	err = getDB.Exec("DELETE FROM public.clients WHERE clients.phone_num = ?", client.PhoneNum)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord Order) IsValid() error {
	if len(ord.PhoneNum) == 0 {
		return errors.New("fields phone_num are required")
	}
	return nil
}

// Show SELECT * FROM orders WHERE client_id = ord.Client_id
func (ord *Order) Show() error {
	getDB := GetDB()
	err := getDB.Where("client_phone_num = ?", ord.PhoneNum).First(&ord)
	if err != nil {
		return err.Error
	}

	return nil
}

func (ord *Order) Add() error {
	err := ord.IsValid()
	if err != nil {
		return err
	}

	getDB := GetDB()
	if err := getDB.Create(&ord); err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Update() error {
	err := ord.IsValid()
	if err != nil {
		return err
	}

	getDB := GetDB()
	if err := getDB.Model(&ord).Where("client_phone_num = ?", ord.PhoneNum).Updates(&ord); err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Del() error {
	getDB := GetDB()
	err := getDB.Delete(&ord, ord.PhoneNum)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
