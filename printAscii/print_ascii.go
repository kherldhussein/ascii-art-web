package webAscii

import (
	"net/http"
	"strings"

	send "webAscii/utils"
)

func printWord(word string, asciiArtGrid [][]string) string {
	var result strings.Builder
	for i := 1; i <= 8; i++ {
		for _, char := range word {
			index := int(char - 32)
			result.WriteString(asciiArtGrid[index][i])
		}
		result.WriteString("\n")
	}
	return result.String()
}

func PrintArt(w http.ResponseWriter, str string, asciiArtGrid [][]string) string {
	var result strings.Builder

	switch str {
	case "":
		send.SendError(w, "insufficient input", http.StatusBadRequest)
	case "\\n":
		result.WriteString("\n")
	default:
		lines := strings.Split(strings.ReplaceAll(str, "\\n", "\n"), "\n")
		num := 0
		for _, line := range lines {
			for _, ch := range line {
				if ch < ' ' || ch > '~' {
					send.SendError(w, string(ch)+" is a non-printable ASCII character", http.StatusBadRequest)
					return ""
				}
			}
			if line == "" {
				num++
				if num < len(lines) {
					result.WriteString("\n")
				}
			} else {
				result.WriteString(printWord(line, asciiArtGrid))
			}
		}
	}
	return result.String()
}
