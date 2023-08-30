package http

import (
	"encoding/json"
	"net/http"

	"github.com/yugarinn/github-issues-notificator/internal"
)


func createNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var notificationCreationRequest NotificationCreationRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&notificationCreationRequest)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if notificationCreationRequest.RepositoryUrl == "" || notificationCreationRequest.Label == "" || notificationCreationRequest.Email == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	input := internal.CreateNotificationInput{
		RepositoryUrl: notificationCreationRequest.RepositoryUrl,
		Label: notificationCreationRequest.Label,
		Email: notificationCreationRequest.Email,

	}
	notificationCreationResult := internal.CreateNotification(input)

	if notificationCreationResult.Error != nil {
		http.Error(w, "Error on notification creation", http.StatusUnprocessableEntity)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
}
