package memory

import (
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// Добавление маркетов
func (s *Storage) AddMarkets() error {

	err := godotenv.Load("./SpotInstrumentService/config/market/.env") //Подгрузка переменной окружения
	if err != nil {
		return err
	}

	path := os.Getenv("MARKETS")
	file, err := os.ReadFile(path) //Чтение файла с рынками
	if err != nil {
		return fmt.Errorf("ошибка прочтения файла:%w", err)
	}

	re := regexp.MustCompile(`([a-zA-Z0-9]+)+\s?([a-zA-Z0-9]+)?`) //Извление названия рынков и добавление их в хранилище
	markets := re.FindAllString(string(file), -1)
	//Заполнение storage
	for i, m := range markets {
		s.date[m] = &domainmarket.Market{ID: int64(i), Name: m, Enable: true, Delete_at: nil}
	}
	return nil
}
