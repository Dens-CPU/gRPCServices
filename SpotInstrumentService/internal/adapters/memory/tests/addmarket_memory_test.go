package memory_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/DencCPU/gRPCServices/SpotInstrumentService/internal/adapters/memory"
	"go.uber.org/zap"
)

func TestStorage_AddMarkets(t *testing.T) {
	tDir := t.TempDir()
	path := filepath.Join(tDir, "market.txt")
	filedate := "Binance, TradingView, Coinbase Exchange, Upbit"
	err := os.WriteFile(path, []byte(filedate), 0644)
	if err != nil {
		t.Fatal("Ошибка записи в файл:", err)
	}
	tests := []struct {
		name string // description of this test case
		// Named input parameters for receiver constructor.
		logger *zap.Logger
		// Named input parameters for target function.
		path    string
		wantErr bool
	}{
		{
			name:    "add markets to storage",
			path:    path,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, err := memory.NewStorage(tt.logger)
			if err != nil {
				t.Fatalf("could not construct receiver type: %v", err)
			}
			gotErr := s.AddMarkets(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("AddMarkets() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("AddMarkets() succeeded unexpectedly")
			}
		})
	}
}
