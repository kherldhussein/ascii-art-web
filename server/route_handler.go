package webAscii

import (
	"fmt"
	"html/template"
	"net/http"

	send "webAscii/utils"
)

func Handl(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	tmpl, err := template.ParseFiles("./templates/index.html")
	if err != nil {
		send.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		fmt.Printf("error %v", err)
	}
}
