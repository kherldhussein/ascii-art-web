package webAscii

import (
	"fmt"
	"os"
)

func WriteAscii(content, fileName string) error {
	err := os.WriteFile(fileName, []byte(content), 0o644)
	if err != nil {
		return fmt.Errorf("error while creating a file: %v", err)
	}
	return err
}
