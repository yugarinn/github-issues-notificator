package main

import (
	"github.com/yugarinn/github-issues-notificator/core"
	"github.com/yugarinn/github-issues-notificator/http"
	"github.com/yugarinn/github-issues-notificator/worker"
)


func main() {
	app := core.BootstrapApplication()

	go worker.InitiNotificationWorker(app)
	go http.InitServer(app)

	select {}
}
