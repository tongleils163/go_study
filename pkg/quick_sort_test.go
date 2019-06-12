package pkg

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		data  []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.data, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
