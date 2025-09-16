package main

import (
	"encoding/json"
	"net/http"

	"github.com/danilobml/email-sender/cmd/api/mailer"
)

type EmailRequest struct {
	To      []string `json:"to"`
	Subject string   `json:"subject"`
	Body    string   `json:"body"`
}

func handleMail(w http.ResponseWriter, r *http.Request) {
	var requestBody EmailRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Failed to parse", http.StatusBadRequest)
	}

	err = mailer.SendMail(requestBody.To, requestBody.Subject, requestBody.Body)
	if err != nil {
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
	}
}
