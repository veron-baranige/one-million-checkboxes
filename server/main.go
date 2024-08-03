package main

import (
	"log"
	"net/http"

	"github.com/veron-baranige/one-million-checkboxes/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.WebsocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
