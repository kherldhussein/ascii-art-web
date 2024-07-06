package webAscii

import (
	"fmt"
	"net/http"

	check "webAscii/checksum"
	print "webAscii/printAscii"
	output "webAscii/readWrite"
)

var banners = map[string]string{
	"standard":   "public/standard.txt",
	"thinkertoy": "public/thinkertoy.txt",
	"shadow":     "public/shadow.txt",
}

func AsciiServer(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() %v", err), http.StatusBadRequest)
	}

	input := r.FormValue("Input")
	banner := r.FormValue("Banner")

	str := ""

	all := []string{"standard", "thinkertoy", "shadow"}

	if banner == "all" {
		for i, bn := range all {
			if i != 0 {
				str += "\n"
			}
			str += writeAscii(w, bn, input)
		}
	} else {
		str += writeAscii(w, banner, input)
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, str)
}

func writeAscii(w http.ResponseWriter, banner, input string) string {
	filename, ok := banners[banner]
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return "Invalid banner specified\n"
	}

	err := check.ValidateFileChecksum(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error downloading or validating file: %v", err), http.StatusInternalServerError)
		return "Error generating ASCII art"
	}

	asciiArtGrid, err := output.ReadAscii(filename)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading ASCII art: %v", err), http.StatusNoContent)
		return "Error generating ASCII art"
	}

	str := print.PrintArt(w, input, asciiArtGrid)
	return str
}
