package webAscii

import (
	"fmt"
	"net/http"
	"strings"
)

func printWord(w http.ResponseWriter, word string, asciiArtGrid [][]string) string {
	result := ""
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			if index < 0 || index >= len(asciiArtGrid) {
				http.Error(w, "non-printable ascii character", http.StatusMethodNotAllowed)
				return ""
			} else {
				result += asciiArtGrid[index][i]
			}
		}
		result += "\n"
	}
	return result
}

func PrintArt(w http.ResponseWriter, str string, asciiArtGrid [][]string) string {
	result := ""
	switch str {
	case "":
		fmt.Print()
	case "\\n":
		fmt.Println()
	case "\\r", "\\f", "\\v", "\\t", "\\b", "\\a":
		http.Error(w, "Unsupported escape sequence", http.StatusMethodNotAllowed)
		return ""
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
					result += "\n"
					continue
				}
			} else {
				return printWord(w, word, asciiArtGrid)
			}
		}
	}
	return result
}
