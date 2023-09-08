package core

import (
	"cloud.google.com/go/firestore"

	"github.com/yugarinn/github-issues-notificator/database"
	"github.com/yugarinn/github-issues-notificator/lib"
)


type App struct {
	Database *firestore.Client
	GithubClient *lib.GithubClient
	EmailClient *lib.EmailClient
}

func BootstrapApplication() *App {
    app := App{
		Database: database.Firebase(),
		GithubClient: &lib.GithubClient{},
		EmailClient: &lib.EmailClient{},
    }

	return &app
}
