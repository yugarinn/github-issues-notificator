package core

import (
	"cloud.google.com/go/firestore"

	"github.com/yugarinn/github-issues-notificator/database"
)


type App struct {
	Database *firestore.Client
}

func BootstrapApplication() *App {
    app := App{
		Database: database.Firebase(),
    }

	return &app
}
