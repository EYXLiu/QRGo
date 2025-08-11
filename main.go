package main

import (
	"log"
	"net/http"

	"QRGo/handlers"
)

func main() {
	http.HandleFunc("/qrgo", handlers.QRHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
