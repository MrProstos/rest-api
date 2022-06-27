package db

import (
	"errors"
	"fmt"
)

type Tables interface {
	isValid() error
	Insert() (Tables, error)
	Select() (Tables, error)
	Update() (Tables, error)
	Delete() (Tables, error)
}

type client struct {
	PhoneNum  string  `gorm:"primaryKey;unique;type:varchar;not null" json:"phone_num,omitempty"`
	Firstname string  `gorm:"type:varchar;not null" json:"firstname,omitempty"`
	Lastname  string  `gorm:"type:varchar;not null" json:"lastname,omitempty"`
	Birthday  string  `gorm:"type:date;not null" json:"birthday,omitempty"`
	Orders    []order `gorm:"foreignKey:PhoneNum" json:"orders,omitempty"`
}

func (client *client) isValid() error {
	if len(client.Firstname) == 0 || len(client.Lastname) == 0 ||
		len(client.PhoneNum) == 0 || len(client.Birthday) == 0 {
		return errors.New("fields phone_num, firstname, lastname, birthday are required")
	}
	return nil
}

func (client *client) Select() (Tables, error) {
	fmt.Println(client)
	if len(client.PhoneNum) == 0 {
		return nil, errors.New("fields phone_num are required")
	}

	err := GetConn().Model(&client).Where("phone_num = ?", client.PhoneNum).First(&client)
	if err.Error != nil {
		return nil, err.Error
	}

	err = GetConn().Model(&order{}).Where("phone_num = ?", client.PhoneNum).Find(&client.Orders)
	if err.Error != nil {
		return nil, err.Error
	}

	return client, nil
}

func (client *client) Insert() (Tables, error) {
	if err := client.isValid(); err != nil {
		return nil, err
	}

	err := GetConn().Model(&client).Create(&client)
	if err.Error != nil {
		return nil, err.Error
	}

	return client, nil
}

func (client *client) Update() (Tables, error) {
	if len(client.PhoneNum) == 0 {
		return nil, errors.New("fields phone_num are required")
	}

	err := GetConn().Model(&client).Where("phone_num = ?", client.PhoneNum).Updates(&client)
	if err.Error != nil {
		return nil, err.Error
	}

	for _, val := range client.Orders {
		err = GetConn().Model(&val).Where("phone_num = ?", client.PhoneNum).Updates(&val)
		if err.Error != nil {
			return nil, err.Error
		}

	}

	return client, nil
}

func (client *client) Delete() (Tables, error) {
	order := NewOrder()

	err := GetConn().Model(&order).Where("phone_num = ?", client.PhoneNum).Delete(&order)
	if err.Error != nil {
		return nil, err.Error
	}

	err = GetConn().Model(&client).Where("phone_num = ?", client.PhoneNum).Delete(&client)
	if err.Error != nil {
		return nil, err.Error
	}

	return client, nil
}

func NewClient() *client {
	return &client{}
}

type order struct {
	ID       uint   `gorm:"primaryKey"`
	PhoneNum string `gorm:"type:varchar;not null" json:"phone_num,omitempty"`
	To       string `json:"to,omitempty"`
	Body     string `json:"body,omitempty"`
	Status   uint   `gorm:"not null;default:1" json:"status,omitempty"`
}

func (order *order) isValid() error {
	if len(order.PhoneNum) == 0 || len(order.To) == 0 || len(order.Body) == 0 {
		return errors.New("fields phone_num, to, body are required")
	}
	return nil
}

func (order *order) Select() (Tables, error) {
	if len(order.PhoneNum) == 0 {
		return nil, errors.New("fields phone_num are required")
	}

	err := GetConn().Where("phone_num = ?", order.PhoneNum).First(&order)
	if err.Error != nil {
		return nil, err.Error
	}

	return order, nil
}

func (order *order) Insert() (Tables, error) {
	if err := order.isValid(); err != nil {
		return nil, err
	}

	err := GetConn().Model(&order).Create(&order)
	if err.Error != nil {
		return nil, err.Error
	}
	return order, nil
}

func (order *order) Update() (Tables, error) {
	if len(order.PhoneNum) == 0 {
		return nil, errors.New("fields phone_num are required")
	}

	err := GetConn().Model(&order).Where("phone_num = ?", order.PhoneNum).Updates(&order)
	if err.Error != nil {
		return nil, err.Error
	}
	return order, nil
}

func (order *order) Delete() (Tables, error) {
	err := GetConn().Model(&order).Where("phone_num = ?", order.PhoneNum).Delete(&order)
	if err.Error != nil {
		return nil, err.Error
	}
	return order, nil
}

func NewOrder() *order {
	return &order{}
}
