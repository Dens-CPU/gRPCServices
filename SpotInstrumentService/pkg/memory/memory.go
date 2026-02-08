// В данном файле прописаны используемые структуры
package memory

import "time"

// Структура Market
type Market struct {
	ID        int64
	Name      string
	Enable    bool
	Delete_at *time.Time
}

// Хранилище данных о рынках
type Storage map[string]*Market

func NewStorage(size int) Storage {
	return make(Storage)
}
