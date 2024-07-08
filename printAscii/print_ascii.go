package webAscii

import (
	"net/http"
	"strings"
)

func printWord(word string, asciiArtGrid [][]string) string {
	result := ""
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			result += asciiArtGrid[index][i]
		}
		result += "\n"
	}
	return result
}

func PrintArt(w http.ResponseWriter, str string, asciiArtGrid [][]string) string {
	result := ""
	switch str {
	case "":
		http.Error(w, "400: Bad request", http.StatusBadRequest)
	case "\\n":
		result += "\n"
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		words := strings.Split(s, "\n")
		num := 0
		for _, word := range words {
			for _, ch := range word {
				if ch < ' ' || ch > '~' {
					http.Error(w, string(ch)+" is non-printable ascii character", http.StatusBadRequest)
					return ""
				}
			}
			if word == "" {
				num++
				if num < len(words) {
					result += "\n"
					continue
				}
			} else {
				result += printWord(word, asciiArtGrid)
			}
		}
	}
	return result
}
