package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	config := LoadConfig()

	sheetsService, err := NewSheetsService(config, "credentials.json")
	if err != nil {
		log.Fatalf("Failed to initialize Google Sheets API: %v", err)
	}

	whatsappHandler := NewWhatsAppHandler(config, sheetsService)

	r := chi.NewRouter()
	r.Get("/webhook", whatsappHandler.VerifyWebhook)
	r.Post("/webhook", whatsappHandler.HandleMessage)

	log.Printf("Server running on port %s", config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, r))
}
