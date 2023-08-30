package main

import (
	"github.com/yugarinn/github-issues-notificator/core"
	"github.com/yugarinn/github-issues-notificator/worker"
	"github.com/yugarinn/github-issues-notificator/http"
)


func main() {
	app := core.BootstrapApplication()
	defer app.Database.Close()

	worker.InitiNotificationWorker(app)
	http.InitServer(app)
}
