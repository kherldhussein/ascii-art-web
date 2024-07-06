package output

import (
	"testing"

	output "webAscii/printAscii"
	in "webAscii/readWrite"
)

func TestWriteArt(t *testing.T) {
	asciiArtGrid, err := in.ReadAscii("assets/standard.txt")
	if err != nil {
		t.Errorf("Error Reading: %v", err)
		return
	}

	testCases := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "Empty input",
			input:    "",
			expected: "",
			wantErr:  false,
		},

		{
			name:     "Single word",
			input:    "Hello",
			expected: " _    _          _   _          \n| |  | |        | | | |         \n| |__| |   ___  | | | |   ___   \n|  __  |  / _ \\ | | | |  / _ \\  \n| |  | | |  __/ | | | | | (_) | \n|_|  |_|  \\___| |_| |_|  \\___/  \n                                \n                                \n",
			wantErr:  false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := output.WriteArt(tc.input, asciiArtGrid)
			if (err != nil) != tc.wantErr {
				t.Errorf("WriteArt(%q) error = %v, wantErr %v", tc.input, err, tc.wantErr)
				return
			}
			if result != tc.expected {
				t.Errorf("WriteArt(%q) = %q, want %q", tc.input, result, tc.expected)
			}
		})
	}
}
