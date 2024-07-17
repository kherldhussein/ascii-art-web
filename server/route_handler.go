package webAscii

import (
	"fmt"
	"html/template"
	"net/http"

	send "webAscii/utils"
)

// Handl handles HTTP requests at the root path ("/").
// Serves the index.html template located in the "./templates/" directory.
func Handl(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		send.SendError(w, "Error 404: PAGE NOT FOUND", http.StatusNotFound)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		fmt.Printf("error %v", err)
	}
}
