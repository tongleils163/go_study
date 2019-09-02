package pkg

import (
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name string
	}{
		{name: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Parse()
		})
	}
}

func Test_test11(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			test11(tt.args.s)
		})
	}
}
