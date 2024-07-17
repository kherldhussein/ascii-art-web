package webAscii

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"

	send "webAscii/utils"
)

var files = map[string]bool{
	"public/shadow.txt":     true,
	"public/standard.txt":   true,
	"public/thinkertoy.txt": true,
}

func ValidateFileName(file string) bool {
	_, ok := files[file]
	return ok
}

func ReadAscii(filename string, w http.ResponseWriter) ([][]string, error) {
	if !ValidateFileName(filename) {
		send.SendError(w, "Error 500: Internal server error", http.StatusInternalServerError)
		return nil, fmt.Errorf("unsupported file name: %s", filename)
	}

	if !strings.HasSuffix(filename, ".txt") {
		send.SendError(w, "Error 500: Internal server error", http.StatusInternalServerError)
		return nil, fmt.Errorf("unsupported file format: %s", filename)
	}

	file, err := os.Open(filename)
	if err != nil {
		send.SendError(w, fmt.Sprintf("Error 404 Not Found: %v", err), http.StatusNotFound)
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	defer file.Close()
	var (
		asciiArtGrid [][]string
		asciLine     []string
	)

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		lines := scanner.Text()
		asciLine = append(asciLine, lines)
		count++
		if count == 9 {
			asciiArtGrid = append(asciiArtGrid, asciLine)
			count = 0
			asciLine = []string{}
		}
	}

	if err := scanner.Err(); err != nil {
		send.SendError(w, "Error 500: Internal server error", http.StatusInternalServerError)
		return nil, fmt.Errorf("error scanning file: %w", err)
	}
	return asciiArtGrid, nil
}
