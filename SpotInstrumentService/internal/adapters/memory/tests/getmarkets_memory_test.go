package memory_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/DencCPU/gRPCServices/Shared/logger"
	"github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/adapters/memory"
	domainusers "github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/domain/users"
)

func TestStorage_GetEnableMarkets(t *testing.T) {
	logger, _ := logger.NewLogger()

	tDir := t.TempDir()
	path := filepath.Join(tDir, "market.txt")
	filedate := "Binance, TradingView, Coinbase Exchange, Upbit"
	err := os.WriteFile(path, []byte(filedate), 0644)
	if err != nil {
		t.Fatal("write to file error:", err)
	}

	s, err := memory.NewStorage(logger)
	if err != nil {
		t.Fatal("initialization error:", err)
	}

	err = s.AddMarkets(path)
	if err != nil {
		t.Fatal("error add markets to storage", err)
	}
	input := domainusers.Input{
		UserRole:  domainusers.USER_ROLE_BASIC_USER,
		PageSize:  0,
		PageToken: "",
	}
	fmt.Println("UserRole:", input.UserRole)

	date, pagetoken := s.GetEnableMarkets(input)

	if len(date) == 0 {
		t.Fatal("Enable markets list is empty")
	}

	fmt.Println("quatity enable markets:", len(date))

	if pagetoken == "" {
		t.Fatal("pageToken is empty")
	}

	re := regexp.MustCompile(`([a-zA-Z0-9]+)+\s?([a-zA-Z0-9]+)?`)
	markets := re.FindAllString(filedate, -1)

	var checkMap = make(map[string]struct{})
	for _, m := range markets {
		checkMap[m] = struct{}{}
	}
	for _, get := range date {
		if _, exsist := checkMap[get.Name]; !exsist {
			log.Fatal("error.there are no values ​​in the resulting list:", err)
		}
	}

}
