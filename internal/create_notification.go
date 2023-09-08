package internal

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/yugarinn/github-issues-notificator/database"
)

type NotificationFilters struct {
	Author		string
	Assignee	string
	Label 		string
	Title 		string
}

type CreateNotificationInput struct {
	RepositoryUri string
	Filters       NotificationFilters
	Email         string
}

type CreateNotificationResult struct {
	Success      bool
	Error        error
}

func CreateNotification(input CreateNotificationInput) CreateNotificationResult {
	if repositoryExists(input.RepositoryUri) == false {
		return CreateNotificationResult{
			Success: false,
			Error: errors.New("provided_repository_not_found"),
		}
	}

	context := context.Background()
	firebase := database.Firebase()

	now := time.Now()

	notification := map[string]interface{}{
		"repositoryUri": 	input.RepositoryUri,
		"email": 			input.Email,
		"filters": 			input.Filters,
		"confirmationCode":	generateNotificationConfirmationCode(&input),
		"isConfirmed": 		false,
		"createdAt":        now,
		"updatedAt":        now,
	}

	_, error := firebase.Collection("notifications").NewDoc().Create(context, notification)

	return CreateNotificationResult{
		Success: error != nil,
		Error: error,
	}
}

func repositoryExists(repositoryUri string) bool {
	url := fmt.Sprintf("https://github.com%s", repositoryUri)
	response, err := http.Get(url)

	if response.StatusCode != 200 || err != nil {
		return false
	}

	return true
}

func generateNotificationConfirmationCode(input *CreateNotificationInput) string {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	baseString := hex.EncodeToString(bytes)
	baseCode := fmt.Sprintf("%s-%s-%s", input.RepositoryUri, input.Email, baseString)

	hash := sha256.New()
	hash.Write([]byte(baseCode))
	hashedBytes := hash.Sum(nil)

	confirmationCode := hex.EncodeToString(hashedBytes)

	return confirmationCode
}
