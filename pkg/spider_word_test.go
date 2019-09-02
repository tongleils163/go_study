package pkg

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestProcess(t *testing.T) {
	tests := []struct {
		name string
	}{{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Process()
		})
	}
}
