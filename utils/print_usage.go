package webAscii

import (
	"fmt"
	"net/http"
)

func PrintUsage() {
	fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}

func SendError(w http.ResponseWriter, message string, status int) {
	http.Error(w, message, status)
}
