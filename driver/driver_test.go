package driver

import (
	"github.com/joho/godotenv"
	"testing"

	_ "github.com/lib/pq"
)

func init() {
	godotenv.Load("test.env")
}

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: "Test DB Connection"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ConnectDB()
		})
	}
}
