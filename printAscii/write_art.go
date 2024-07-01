package webAscii

import (
	"fmt"
	"strings"
)

// Generates ASCII art from the input string using the provided ASCII art grid.
func WriteArt(str string, asciiArtGrid [][]string) (string, error) {
	var result strings.Builder

	switch str {
	case "":
		result.WriteString("")
	case "\\n":
		result.WriteString("\n")
	case "\\r", "\\f", "\\v", "\\t", "\\b", "\\a":
		return "", fmt.Errorf("error: unsupported escape sequence '%s'", str)
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		s = strings.ReplaceAll(s, "\\r", "\r")
		s = strings.ReplaceAll(s, "\\f", "\f")
		s = strings.ReplaceAll(s, "\\v", "\v")
		s = strings.ReplaceAll(s, "\\t", "\t")
		s = strings.ReplaceAll(s, "\\b", "\b")
		s = strings.ReplaceAll(s, "\\a", "\a")
		words := strings.Split(s, "\n")
		num := 0
		for _, word := range words {
			if word == "" {
				num++
				if num < len(words) {
					result.WriteString("\n")
					continue
				}
			} else {
				err := writeWord(word, asciiArtGrid, &result)
				if err != nil {
					return "", err
				}
			}
		}
	}
	return result.String(), nil
}

// Writes the ASCII representation of a word using the provided ASCII art grid.
func writeWord(word string, asciiArtGrid [][]string, res *strings.Builder) error {
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			if index < 0 || index >= len(asciiArtGrid) {
				return fmt.Errorf("unknown character: %q", char)
			} else {
				res.WriteString(asciiArtGrid[index][i])
			}
		}
		res.WriteString("\n")
	}
	return nil
}
