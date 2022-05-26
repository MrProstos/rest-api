package db

type ManageDb interface {
	IsValid() error //Проверка на валидность данных
	Show() error    //Получить данные клиента по номеру телефона
	Add() error     //Добавляет клиента в базу данных
	Update() error  //Обновление данных клиента
	Del() error     //Удаление клиента
}
