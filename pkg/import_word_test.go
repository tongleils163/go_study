package pkg

import "testing"

func TestImportWord(t *testing.T) {
	initMysqlEngine()
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{path: "./static/审完的"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ImportWord(tt.args.path)
		})
	}
}
