package domainmarket

import (
	"time"

	domainusers "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/users"
)

// Структура хранения данных о рынке
type Market struct {
	ID         string
	Name       string
	Enable     bool
	DeleteAt   *time.Time
	UserAccess domainusers.UserRole
}

// Конструктор для создания нового рынка
func NewMarket(id, name string) *Market {
	m := Market{
		ID:       id,
		Name:     name,
		Enable:   true,
		DeleteAt: nil,
	}
	return &m
}
