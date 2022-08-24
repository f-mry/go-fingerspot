package main

import (
	"log"
	"net/http"

	"github.com/farhanmry/go-fingerspot"
)

func main() {
	server := fingerspot.NewWebhookServer()
	log.Fatal(http.ListenAndServe(":42069", server.Server()))
}
