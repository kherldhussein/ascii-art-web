package webAscii

import (
	"fmt"
	"net/http"

	check "webAscii/checksum"
	print "webAscii/printAscii"
	output "webAscii/readWrite"
	send "webAscii/utils"
)

var banners = map[string]string{
	"standard":   "public/standard.txt",
	"thinkertoy": "public/thinkertoy.txt",
	"shadow":     "public/shadow.txt",
}

func AsciiServer(w http.ResponseWriter, r *http.Request) {
	// chech method
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Parse for data from the client
	if err := r.ParseForm(); err != nil {
		send.SendError(w, fmt.Sprintf("ParseForm() %v", err), http.StatusBadRequest)
		return
	}
	// retrieve value associated with the
	text := r.FormValue("Text")
	banner := r.FormValue("Banner")
	for param := range r.Form {
		if param != "Text" && param != "Banner" {
			send.SendError(w, "Error 400: Bad request", http.StatusBadRequest)
			break
		}
	}

	if banner == "" || text == "" {
		send.SendError(w, "Error 400 Bad request: nothing is specified", http.StatusBadRequest)
		return
	}
	str := ""

	all := []string{"standard", "thinkertoy", "shadow"}

	if banner == "all" {
		for i, bn := range all {
			if i != 0 {
				str += "\n"
			}
			str += writeAscii(w, bn, text)
		}
	} else {
		str += writeAscii(w, banner, text)
	}

	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, str)
}

func writeAscii(w http.ResponseWriter, banner, text string) string {
	filename, ok := banners[banner]
	if !ok {
		send.SendError(w, "Error 404: Not Found: Invalid banner specified\n", http.StatusNotFound)
		return ""
	}

	if err := check.ValidateFileChecksum(w, filename); err != nil {
		send.SendError(w, fmt.Sprintf("Error 404: Error processing file: %v", err), http.StatusNotFound)
		return ""
	}

	asciiArtGrid, err := output.ReadAscii(filename, w)
	if err != nil {
		send.SendError(w, fmt.Sprintf("Error 500: Internal Server Error: Error reading ASCII art:%v", err), http.StatusInternalServerError)
		return ""
	}

	return print.PrintArt(w, text, asciiArtGrid)
}
