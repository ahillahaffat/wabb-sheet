package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

type WhatsAppHandler struct {
	config *Config
	sheets *SheetsService
}

func NewWhatsAppHandler(config *Config, sheets *SheetsService) *WhatsAppHandler {
	return &WhatsAppHandler{
		config: config,
		sheets: sheets,
	}
}

func (h *WhatsAppHandler) VerifyWebhook(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token == h.config.VerifyToken {
		fmt.Fprint(w, challenge)
		log.Println("Webhook Verified")
	} else {
		http.Error(w, "Forbidden", http.StatusForbidden)
	}
}

func (h *WhatsAppHandler) HandleMessage(w http.ResponseWriter, r *http.Request) {
	var body map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	bodyJSON, _ := json.MarshalIndent(body, "", "  ")
	log.Println("Webhook Payload Received:\n", string(bodyJSON))

	message, from, err := h.parseWebhookPayload(body)
	if err != nil {
		log.Printf("Error parsing webhook payload: %v", err)
		w.WriteHeader(http.StatusOK)
		return
	}

	log.Printf("Incoming message: %s", message)
	log.Printf("Sender: %s", from)

	if err := h.sheets.SaveMessage(message); err != nil {
		log.Printf("Failed to save to Google Sheet: %v", err)
	} else {
		if err := h.sheets.StyleHeader(); err != nil {
			log.Printf("Failed to style header: %v", err)
		}
	}

	if err := h.sendConfirmationMessage(from); err != nil {
		log.Printf("Failed to send confirmation message: %v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *WhatsAppHandler) parseWebhookPayload(body map[string]interface{}) (string, string, error) {
	entry, ok := body["entry"].([]interface{})
	if !ok || len(entry) == 0 {
		return "", "", fmt.Errorf("invalid entry format")
	}

	entryData, ok := entry[0].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid entry data format")
	}

	changes, ok := entryData["changes"].([]interface{})
	if !ok || len(changes) == 0 {
		return "", "", fmt.Errorf("invalid changes format")
	}

	changeData, ok := changes[0].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid change data format")
	}

	value, ok := changeData["value"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid value format")
	}

	messages, ok := value["messages"].([]interface{})
	if !ok || len(messages) == 0 {
		return "", "", fmt.Errorf("no messages found")
	}

	messageData, ok := messages[0].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid message data format")
	}

	from, ok := messageData["from"].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid from format")
	}

	textData, ok := messageData["text"].(map[string]interface{})
	if !ok {
		return "", "", fmt.Errorf("invalid text format")
	}

	body_text, ok := textData["body"].(string)
	if !ok {
		return "", "", fmt.Errorf("invalid body format")
	}

	return body_text, from, nil
}

func (h *WhatsAppHandler) sendConfirmationMessage(to string) error {
	client := resty.New()

	url := fmt.Sprintf("https://graph.facebook.com/v19.0/%s/messages", h.config.PhoneNumberID)

	resp, err := client.R().
		SetHeader("Authorization", "Bearer "+h.config.AccessToken).
		SetHeader("Content-Type", "application/json").
		SetBody(map[string]interface{}{
			"messaging_product": "whatsapp",
			"to":                to,
			"text": map[string]string{
				"body": "✅ Laporan kamu sudah dicatat di Google Spreadsheet.",
			},
		}).
		Post(url)

	if err != nil {
		return fmt.Errorf("failed to send message: %v", err)
	}

	log.Printf("Response Status Code: %d", resp.StatusCode())
	log.Printf("Response Body: %s", resp.String())

	return nil
}
