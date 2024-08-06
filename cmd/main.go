package main

import (
	"log"
	"net/http"

	"github.com/my-golang-webhooks/internal/handlers"
)

func main() {
	// Initialize the webhook handler
	webhookHandler := handlers.NewWebhookHandler()

	// Set up the routes for handling webhooks
	http.HandleFunc("/webhook", webhookHandler.HandleWebhook)

	// Start the web server
	log.Println("Starting server on port 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}