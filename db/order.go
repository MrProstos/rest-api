package db

import (
	"errors"
)

type Order struct {
	Client_id uint
	Title     string
	To        string
	Body      string
	Status    uint
}

func (ord Order) IsValid() error {
	if ord.Client_id == 0 && ord.Status == 0 {
		return errors.New("fields Client_id, Status are required")
	}
	return nil
}

func (ord Order) Add() error {
	if err := ord.IsValid(); err != nil {
		return err
	}
	db := GetDB().Create(&ord)
	if db.Error != nil {
		return db.Error
	}
	return nil
}

func (ord Order) Update() error {
	err := GetDB().Model(&ord).Update(ord)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Del() error {
	db := GetDB().Delete(&ord, 1)
	if db.Error != nil {
		return db.Error
	}
	return nil
}
