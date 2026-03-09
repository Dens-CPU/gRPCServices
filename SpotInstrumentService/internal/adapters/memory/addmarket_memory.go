package memory

import (
	configfile "Academy/gRPCServices/Shared/config"
	domainmarket "Academy/gRPCServices/SpotInstrumentService/internal/domain/market"
	"io"
	"regexp"
)

// Добавление маркетов
func (s *Storage) AddMarkets() error {

	file, err := configfile.NewConfigFile("./SpotInstrumentService/config/market/.env", "MARKETS")
	if err != nil {
		return err
	}

	str, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`([a-zA-Z0-9]+)+\s?([a-zA-Z0-9]+)?`) //Извление названия рынков и добавление их в хранилище
	markets := re.FindAllString(string(str), -1)
	//Заполнение storage
	for i, m := range markets {
		s.date[m] = &domainmarket.Market{ID: int64(i), Name: m, Enable: true, Delete_at: nil}
	}
	return nil
}
