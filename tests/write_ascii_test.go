package tests

import (
	"os"
	"testing"

	output "webAscii/readWrite"
)

func TestWriteAscii(t *testing.T) {
	content := "Hello, world!"
	fileName := "testfile.txt"

	err := output.WriteAscii(content, fileName)
	if err != nil {
		t.Errorf("WriteAscii(%q, %q) returned error: %v", content, fileName, err)
	}

	data, err := os.ReadFile(fileName)
	if err != nil {
		t.Errorf("Failed to read file %q: %v", fileName, err)
	}
	if string(data) != content {
		t.Errorf("File content mismatch. Expected: %q, Got: %q", content, string(data))
	}

	err = os.Remove(fileName)
	if err != nil {
		t.Errorf("Failed to delete test file %q: %v", fileName, err)
	}

	invalidFileName := "testfile.csv"

	err = output.WriteAscii(content, invalidFileName)
	if err == nil {
		t.Errorf("WriteAscii(%q, %q) expected error, but got nil", content, invalidFileName)
	}

	expectedErrorMsg := "usage: go run . --output=<fileName.txt> something standard"
	if err.Error() != expectedErrorMsg {
		t.Errorf("Expected error message %q, got %q", expectedErrorMsg, err.Error())
	}
}
