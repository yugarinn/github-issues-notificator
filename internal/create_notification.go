package internal

import (
	"github.com/yugarinn/github-issues-notificator/database"
)


type CreateNotificationInput struct {
	RepositoryUrl string
	Label         string
	Email         string
}

type CreateNotificationResult struct {
	Notification string
	Success      bool
	Error        error
}

func CreateNotification(input CreateNotificationInput) CreateNotificationResult {
	firebase, _ := database.Firebase()

	firebase.Create()

	return CreateNotificationResult{
		Notification: "notification",
		Success: true,
		Error: nil,
	}
}
