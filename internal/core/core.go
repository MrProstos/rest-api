package core

type Operator struct {
	Username string
	Password string
}

// Client Структура данных Клиент содержащая поля такие же как и в базе данных
type Client struct {
	ID        uint   `gorm:"primaryKey"` // Первичный ключ autoincrement . ПОМЕНЯТЬ НА НОМЕР ТЕЛЕФОНА
	Phone_num string `gorm:"unique;type:varchar;not null" json:"phone_num"`
	Firstname string `gorm:"type:varchar;not null" json:"firstname"`
	Lastname  string `gorm:"type:varchar;not null" json:"lastname"`
	Birthday  string `gorm:"type:varchar;not null" json:"birthday"` // Поменять на тип данных Date
	//OrderID   uint
	Orders []Order
}

type Order struct {
	Client_id uint `gorm:"primaryKey"`
	Title     string
	To        string
	Body      string
	Status    uint
}
