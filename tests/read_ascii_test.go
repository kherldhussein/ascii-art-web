package tests

import (
	"testing"

	output "webAscii/readWrite"
)

func TestReadAscii(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     bool
	}{
		{
			name:     "valid file name",
			filename: "shadow.txt",
			want:     false,
		},
		{
			name:     "invalid file name",
			filename: "invalid.txt",
			want:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := output.ValidateFileName(tt.filename)
			if got != tt.want {
				t.Errorf("validateFileName(%s) = %v, want %v", tt.filename, got, tt.want)
			}
		})
	}
}
