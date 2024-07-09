package main

import (
	"log"
	"net/http"

	server "webAscii/server"
)

func main() {
	http.HandleFunc("/", server.Handl)
	http.HandleFunc("/ascii-art", server.AsciiServer)

	log.Println("Server listening on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Error Running Server: %v ", err)
	}
}
