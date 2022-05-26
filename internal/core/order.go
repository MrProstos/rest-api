package core

import (
	"errors"
	"github.com/MrProstos/rest-api/internal/db"
	"log"
)

type Order struct {
	Client_id uint `gorm:"primaryKey"`
	Title     string
	To        string
	Body      string
	Status    uint
}

func init() {
	getDB := db.GetDB()
	err := getDB.AutoMigrate(&Client{}, &Order{})
	if err != nil {
		log.Fatal(err)
	}
}

func (ord Order) IsValid() error {
	if ord.Status == 0 {
		return errors.New("fields Client_id, Status are required")
	}
	return nil
}

// Show SELECT * FROM orders WHERE client_id = ord.Client_id
func (ord *Order) Show() error {
	getDB := db.GetDB()
	err := getDB.Where("client_id = ?", ord.Client_id).First(&ord)
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

	getDB := db.GetDB()
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

	getDB := db.GetDB()
	if err := getDB.Model(&ord).Updates(&ord); err.Error != nil {
		return err.Error
	}
	return nil
}

func (ord *Order) Del() error {
	getDB := db.GetDB()
	err := getDB.Delete(&ord, 1)
	if err.Error != nil {
		return err.Error
	}
	return nil
}
