package pkg

import "testing"

func TestShuffle(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name:""},
	}
	for _, tt := range tests {
		for true {

			t.Run(tt.name, func(t *testing.T) {
				Shuffle()
			})
		}
	}
}
