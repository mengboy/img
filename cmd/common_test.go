package cmd

import (
	"testing"
)

func TestGetFileNameWithoutExt(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{filePath: "/Users/vector/dev/go/src/github.com/mengboy/img/test.png"}, want: "test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetFileNameWithoutExt(tt.args.filePath); got != tt.want {
				t.Errorf("GetFileNameWithoutExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
