package domainmarket

import "time"

// Структура хранения данных о рынке
type Market struct {
	ID        int64      //Идентификатор маркета
	Name      string     //Название маркета
	Enable    bool       //Доступ к маркету
	Delete_at *time.Time //Время удаления маркета
}

// Конструктор для создания нового рынка
func NewMarket(id int64, name string) *Market {
	m := Market{
		ID:        id,
		Name:      name,
		Enable:    true,
		Delete_at: nil,
	}
	return &m
}
