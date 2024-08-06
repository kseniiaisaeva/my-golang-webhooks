package handlers

import (
	"fmt"
	"net/http"
)

// HandleWebhook processes incoming webhook requests
func HandleWebhook(w http.ResponseWriter, r *http.Request) {
	// Add your webhook handling logic here
	// You can access the request body and headers using r.Body and r.Header

	// Example: Print the request body
	body := make([]byte, r.ContentLength)
	_, err := r.Body.Read(body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	fmt.Println(string(body))

	// Example: Send a response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Webhook received successfully"))
}