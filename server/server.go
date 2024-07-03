package webAscii

import (
	"fmt"
	"log"
	"net/http"

	check "webAscii/checksum"
	print "webAscii/printAscii"
	output "webAscii/readWrite"
)

var banners = map[string]string{
	"standard":   "standard.txt",
	"thinkertoy": "thinkertoy.txt",
	"shadow":     "shadow.txt",
	
}

// Enable CORS middleware
// func enableCors(w *http.ResponseWriter) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

func AsciiServer(w http.ResponseWriter, r *http.Request) {
	// if the method is post
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// check errors
	if err := r.ParseForm(); err != nil {
		http.Error(w, fmt.Sprintf("ParseForm() %v", err), http.StatusBadRequest)
	}

	// grab th einputs
	input := r.FormValue("Input")
	banner := r.FormValue("Banner")

	str := ""

	all := []string {"standard", "thinkertoy", "shadow" }

	if banner == "all"{
		for i, bn := range all{
			if i != 0{
				str += "\n"
			}
			str += writeAscii(bn, input)
		}
	}else{
		str += writeAscii(banner, input)
	}

	

	// write to the
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprint(w, str)
}

func writeAscii(banner, input string)string{
	filename, ok := banners[banner]
	if !ok {
		return "Invalid banner specified"
		return ""
	}

	err := check.ValidateFileChecksum(filename)
	if err != nil {
		log.Printf("Error downloading or validating file: %v", err)
		return "Error generating ASCII art"
		return ""
	}

	asciiArtGrid, err := output.ReadAscii(filename)
	if err != nil {
		log.Fatalf("Error reading ASCII map: %v", err)
		return "Error generating ASCII art"
	}

	str := print.PrintArt(input, asciiArtGrid)
	return str
}