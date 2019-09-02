package pkg

import "testing"

func TestCallJs(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{args: args{word: "aardwolf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CallJs(tt.args.word); got != tt.want {
				t.Errorf("CallJs() = %v, want %v", got, tt.want)
			}
		})
	}
}
