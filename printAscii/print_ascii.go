package webAscii

import (
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

func PrintArt(str string, asciiArtGrid [][]string) string {
	result := ""
	switch str {
	case "":
		result += ""
	case "\\n":
		result += "\n"
	default:
		s := strings.ReplaceAll(str, "\\n", "\n")
		words := strings.Split(s, "\n")
		num := 0
		for _, word := range words {
			for _, ch := range word {
				if ch < ' ' || ch > '~' {
					return string(ch) + " is non-printable ascii character"
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
