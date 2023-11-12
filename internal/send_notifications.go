package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/yugarinn/github-issues-notificator/core"
)


func SendNotifications(app *core.App) {
	ctx := context.Background()
	collection := app.Database.Collection("notifications")

	entries, err := collection.Find(ctx, bson.M{"isConfirmed": true})

	if err != nil {
		// TODO: add proper logging
		fmt.Println("Error retrieving notifications:", err)
		return
	}
	defer entries.Close(ctx)

	for entries.Next(ctx) {
		var notification Notification

		err := entries.Decode(&notification)
		if err != nil {
			// TODO: add proper logging
			fmt.Println("Error decoding notification:", err)
			continue
		}

		issues, _ := retrieveIssuesFor(notification)

		for _, issue := range issues {
			err := sendIssueEmailAlertTo(issue, notification)

			if (err != nil) {
				// TODO: add proper logging
				fmt.Println(err)
			} else {
				updateNotificationLastCheckAt(app, notification)
			}
		}
	}
}

func retrieveIssuesFor(notification Notification) ([]GithubIssue, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := getIssuesUrlFor(notification)
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Non-OK HTTP status: %s\n", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var issues []GithubIssue
	if err := json.Unmarshal(body, &issues); err != nil {
		return nil, err
	}

	return issues, nil
}

func getIssuesUrlFor(notification Notification) string {
	issueUrl := fmt.Sprintf("https://api.github.com/repos/%s/issues?labels=%s", notification.RepositoryUri, url.QueryEscape(notification.Filters.Label))

	if !notification.LastCheckAt.IsZero() {
		formattedTime := notification.LastCheckAt.Format(time.RFC3339)
		issueUrl += "&since=" + url.QueryEscape(formattedTime)
	}

	return issueUrl
}

func sendIssueEmailAlertTo(issue GithubIssue, notification Notification) error {
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	smtpFrom := "github.issues.notificator@gmail.com"
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	smtpAuth := smtp.PlainAuth("", smtpFrom, smtpPassword, smtpHost)

	headers := map[string]string{
		"From": smtpFrom,
		"To": notification.Email,
		"Subject": fmt.Sprintf("New Issue in %s: %s", notification.RepositoryUri, issue.Title),
		"MIME-Version": "1.0",
		"Content-Type": "text/html; charset=UTF-8",
	}

	header := ""
	for k, v := range headers {
		header += fmt.Sprintf("%s: %s\r\n", k, v)
	}

	body := header + "\r\n" + buildAlertEmailBody(issue)

	err := smtp.SendMail(smtpHost + ":" + smtpPort, smtpAuth, smtpFrom, []string{notification.Email}, []byte(body))

	if err != nil {
		// TODO: add proper logging
		return err
	}

	return nil
}

func buildAlertEmailBody(issue GithubIssue) string {
	return fmt.Sprintf("New issue in %s<br/><br/><a href=\"%s\" target=\"_blank\">%s</a>", issue.RepositoryUrl, issue.Url, issue.Title)
}

func updateNotificationLastCheckAt(app *core.App, notification Notification) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := app.Database.Collection("notifications")

	notificationID, _ := primitive.ObjectIDFromHex(notification.ID)
	filter := bson.M{"_id": notificationID}
	update := bson.M{"$set": bson.M{"lastCheckAt": time.Now()}}

	_, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		// TODO: add proper logging
		return
	}
}
