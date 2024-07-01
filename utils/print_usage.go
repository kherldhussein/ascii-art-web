package webAscii

import "fmt"

func PrintUsage() {
	fmt.Print("Usage: go run . [OPTION] [STRING] [BANNER]\n\n")
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
}
