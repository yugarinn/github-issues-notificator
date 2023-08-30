package internal

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/yugarinn/github-issues-notificator/database"
)


type CreateNotificationInput struct {
	RepositoryUrl string
	Label         string
	Email         string
}

type CreateNotificationResult struct {
	Success      bool
	Error        error
}

func CreateNotification(input CreateNotificationInput) CreateNotificationResult {
	context := context.Background()
	firebase := database.Firebase()

	now := time.Now()

	notification := map[string]interface{}{
		"repositoryUrl": 	input.RepositoryUrl,
		"email": 			input.Email,
		"filters": 			fmt.Sprintf("label.%s,", input.Label),
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

func generateNotificationConfirmationCode(input *CreateNotificationInput) string {
	bytes := make([]byte, 16)
	rand.Read(bytes)

	baseString := hex.EncodeToString(bytes)
	baseCode := fmt.Sprintf("%s-%s-%s-%s", input.RepositoryUrl, input.Email, input.Label, baseString)

	hash := sha256.New()
	hash.Write([]byte(baseCode))
	hashedBytes := hash.Sum(nil)

	confirmationCode := hex.EncodeToString(hashedBytes)

	return confirmationCode
}
