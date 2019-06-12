package pkg

import "testing"

func TestProducerCustomer(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name:""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ProducerCustomer()
		})
	}
}
