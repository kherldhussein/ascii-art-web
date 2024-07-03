package main

import (
	"log"
	"net/http"

	server "webAscii/server"
)

func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/ascii", server.AsciiServer)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error Running Server: %v ", err)
	}
}
