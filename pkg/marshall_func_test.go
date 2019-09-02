package pkg

import "testing"

func TestMarshallFunc(t *testing.T) {
	type args struct {
		answer map[int]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "", args: args{answer: map[int]int{
			100: 1,
			101: 5,
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MarshallFunc()
		})
	}
}
