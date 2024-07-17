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

// HTTP handler function that processes ASCII art generation requests.
func AsciiServer(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST. If not, return a 405 Method Not Allowed error.
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data from the request.
	if err := r.ParseForm(); err != nil {
		send.SendError(w, fmt.Sprintf("ParseForm() %v", err), http.StatusBadRequest)
		return
	}

	// Retrieve the values of the "Text" and "Banner" form fields.
	text := r.FormValue("Text")
	banner := r.FormValue("Banner")

	// Check if there are any additional form parameters and return error 400.
	for param := range r.Form {
		if param != "Text" && param != "Banner" {
			send.SendError(w, "Error 400: Bad request", http.StatusBadRequest)
			break
		}
	}

	// If either field is empty, return error 400.
	if banner == "" || text == "" {
		send.SendError(w, "Error 400 Bad request: nothing is specified", http.StatusBadRequest)
		return
	}

	var output string

	// Generate ASCII art for all available banners.
	if banner == "all" {
		for i, bn := range []string{"standard", "thinkertoy", "shadow"} {
			if i != 0 {
				output += "\n"
			}
			output += writeAscii(w, bn, text)
		}
	} else {
		output = writeAscii(w, banner, text)
	}

	// Set the response content type and writes the output to the response.
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, output)
}

// Generates ASCII art for the given banner and text.
func writeAscii(w http.ResponseWriter, banner, text string) string {
	filename, ok := banners[banner]
	if !ok {
		send.SendError(w, "Error 404: Not Found: Invalid banner specified\n", http.StatusNotFound)
		return ""
	}

	// Validate the checksum of the banner file.
	if err := check.ValidateFileChecksum(w, filename); err != nil {
		send.SendError(w, fmt.Sprintf("Error 404: Error processing file: %v", err), http.StatusNotFound)
		return ""
	}

	// Read the ASCII art grid from the banner file.
	asciiArtGrid, err := output.ReadAscii(filename, w)
	if err != nil {
		send.SendError(w, fmt.Sprintf("Error 500: Internal Server Error: Error reading ASCII art:%v", err), http.StatusInternalServerError)
		return ""
	}

	// Print the ASCII art using the provided text.
	return print.PrintArt(w, text, asciiArtGrid)
}
