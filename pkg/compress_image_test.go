package pkg

import (
	"testing"
)

func Test_imageCompress(t *testing.T) {
	type args struct {
		localPath string
		to        string
		Quality   int
		format    string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "", args: args{localPath: "./static/image", to: "./static/image1", Quality: 80, format: "png"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := imageCompress(tt.args.localPath, tt.args.to, tt.args.Quality, tt.args.format); got != tt.want {
				t.Errorf("imageCompress() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSCompressImage(t *testing.T) {
	type args struct {
		localPath string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{localPath: "./static/image"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SCompressImage(tt.args.localPath)
		})
	}
}

func TestGetImgExts(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		wantExt string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExt, err := GetImgExts(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImgExts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotExt != tt.wantExt {
				t.Errorf("GetImgExts() = %v, want %v", gotExt, tt.wantExt)
			}
		})
	}
}

func TestCompressImage(t *testing.T) {
	type args struct {
		path1 string
		path2 string
	}
	tests := []struct {
		name string
		args args
	}{
		{args: args{path1: "./static/image", path2: "./static/jpeg"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			CompressImage(tt.args.path1, tt.args.path2)
		})
	}
}
