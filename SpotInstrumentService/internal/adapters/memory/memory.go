package memory

import (
	"sync"

	domainmarket "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/market"
	"go.uber.org/zap"
)

// In-memory хранилище для хранения рынков
type Storage struct {
	date   map[string]*domainmarket.Market //Хранилище
	keys   []string
	mu     sync.RWMutex
	logger *zap.Logger
}

// Создание нового хранинлища
func NewStorage(logger *zap.Logger) (*Storage, error) {
	s := Storage{
		date:   make(map[string]*domainmarket.Market),
		logger: logger,
	}

	return &s, nil
}
