package memory_test

import (
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/DencCPU/gRPCServices/Shared/logger"
	"github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/adapters/memory"
)

func TestStorage_GetEnableMarkets(t *testing.T) {
	logger, _ := logger.NewLogger()

	tDir := t.TempDir()
	path := filepath.Join(tDir, "market.txt")
	filedate := "Binance, TradingView, Coinbase Exchange, Upbit"
	err := os.WriteFile(path, []byte(filedate), 0644)
	if err != nil {
		t.Fatal("ошибка записи в файл:", err)
	}

	s, err := memory.NewStorage(logger)
	if err != nil {
		t.Fatal("ошибка инициализации хранилища:", err)
	}

	err = s.AddMarkets(path)
	if err != nil {
		t.Fatal("ошибка добафления рынков в хранилище", err)
	}

	date := s.GetEnableMarkets()
	if len(date) == 0 {
		t.Fatal("Список доступных рынков пуст")
	}

	re := regexp.MustCompile(`([a-zA-Z0-9]+)+\s?([a-zA-Z0-9]+)?`)
	markets := re.FindAllString(filedate, -1)

	var checkMap = make(map[string]struct{})
	for _, m := range markets {
		checkMap[m] = struct{}{}
	}
	for _, get := range date {
		if _, exsist := checkMap[get.Name]; !exsist {
			log.Fatal("ошибка.Нет значения в полученном списке:", err)
		}
	}

}
