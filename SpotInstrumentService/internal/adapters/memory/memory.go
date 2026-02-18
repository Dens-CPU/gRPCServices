package memory

import (
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
	"sync"
)

// In-memory хранилище для хранения рынков
type Storage struct {
	date map[string]*domainmarket.Market //Хранилище
	mu   sync.RWMutex
}

// Создание нового хранинлища
func NewStorage() (*Storage, error) {
	s := Storage{
		date: make(map[string]*domainmarket.Market),
	}

	err := s.AddMarkets() //Добавление рынков в хранилище
	if err != nil {
		return nil, err
	}
	return &s, nil
}
