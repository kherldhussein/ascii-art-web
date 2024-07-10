package webAscii

import (
	"fmt"
	"html/template"
	"net/http"

	send "webAscii/utils"
)

func Handl(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		tmpl, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			tmpl, err := template.ParseFiles("./templates/404.html")
			if err != nil {
				send.SendError(w, "Internal Sasdfgdfgerver Error", http.StatusInternalServerError)
				return
			}

			if err = tmpl.Execute(w, nil); err != nil {
				fmt.Printf("error %v", err)
			}
			// send.SendError(w, "Error 404: NOT FOUND", http.StatusNotFound)
			http.NotFound(w, r)
			return
		}

		if err = tmpl.Execute(w, nil); err != nil {
			fmt.Printf("error %v", err)
		}

	}

	tmpl, err := template.ParseFiles("./templates/404.html")
	if err != nil {
		send.SendError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, nil); err != nil {
		fmt.Printf("error %v", err)
	}
	// send.SendError(w, "Error 404: NOT FOUND", http.StatusNotFound)
	http.NotFound(w, r)
	return
}
