package memory

import (
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
	"sync"

	"go.uber.org/zap"
)

// In-memory хранилище для хранения рынков
type Storage struct {
	date   map[string]*domainmarket.Market //Хранилище
	mu     sync.RWMutex
	logger *zap.Logger
}

// Создание нового хранинлища
func NewStorage(logger *zap.Logger) (*Storage, error) {
	s := Storage{
		date:   make(map[string]*domainmarket.Market),
		logger: logger,
	}

	err := s.AddMarkets() //Добавление рынков в хранилище
	if err != nil {
		s.logger.Error("Ошибка добавления рынков",
			zap.Error(err),
		)
		return nil, err
	}
	return &s, nil
}
