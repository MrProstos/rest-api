package db

import (
	"errors"
)

type Order struct {
	Client_id uint `gorm:"primaryKey"`
	Title     string
	To        string
	Body      string
	Status    uint
}

func (ord Order) IsValid() error {
	if ord.Status == 0 {
		return errors.New("fields Client_id, Status are required")
	}
	return nil
}

//SELECT * FROM orders WHERE client_id = ord.Client_id
func (ord *Order) Show() error {
	db := GetDB()
	err := db.Where("client_id = ?", ord.Client_id).First(&ord)
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

	db := GetDB()
	if err := db.Create(&ord); err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Update() error {
	err := ord.IsValid()
	if err != nil {
		return err
	}

	db := GetDB()
	if err := db.Model(&ord).Updates(&ord); err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Del() error {
	db := GetDB()
	err := db.Delete(&ord, 1)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
