package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/yugarinn/github-issues-notificator/core"
)


func SendNotifications(app *core.App) {
	ctx := context.Background()
	collection := app.Database.Collection("notifications")

	entries, err := collection.Find(ctx, bson.M{"isConfirmed": true})

	if err != nil {
		fmt.Println("Error retrieving notifications:", err)
		return
	}
	defer entries.Close(ctx)

	for entries.Next(ctx) {
		var notification Notification

		err := entries.Decode(&notification)
		if err != nil {
			fmt.Println("Error decoding notification:", err)
			continue
		}

		issues, _ := retrieveIssuesFor(notification)

		for _, issue := range issues {
			fmt.Println(issue.Title)
		}
	}

}

func retrieveIssuesFor(notification Notification) ([]GithubIssue, error) {
	var client = &http.Client{Timeout: 10 * time.Second}
	url := fmt.Sprintf("https://api.github.com/repos/%s/issues?labels=%s", notification.RepositoryUri, url.QueryEscape(notification.Filters.Label))
	resp, err := client.Get(url)

	if err != nil {
		return nil, err
	}

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

	resp.Body.Close()

	return issues, nil
}
