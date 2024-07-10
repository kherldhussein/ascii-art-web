package webAscii

import (
	"errors"
	"fmt"
	"os"
	"regexp"
)

func WriteAscii(content, fileName string) error {
	if !isValidFileName(fileName) {
		return errors.New("usage: go run . --output=<fileName.txt> something standard")
	}

	err := os.WriteFile(fileName, []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("error while creating a file: %v", err)
	}
	return err
}

func isValidFileName(fileName string) bool {
	match, err := regexp.MatchString(`^.+\.txt$`, fileName)
	if err != nil {
		fmt.Printf("Error, %v", err)
	}
	return match
}
