package db

type Message struct {
	ID    uint
	Id    string `gorm:"primary_key"` //Принимает Order.Message_ID
	Title string
	To    string
	Body  string
}
