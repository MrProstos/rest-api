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

func (ord *Order) Add() error {
	if err := ord.IsValid(); err != nil {
		return err
	}

	db := GetDB()
	if err := db.Create(&ord); err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Update() error {
	if err := ord.IsValid(); err != nil {
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
	if err := db.Delete(&ord, 1); err.Error != nil {
		return err.Error
	}
	return nil
}
