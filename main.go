package main

import (
	"github.com/yugarinn/github-issues-notificator/worker"
	"github.com/yugarinn/github-issues-notificator/http"
)


func main() {
	worker.InitiNotificationWorker()
	http.InitServer()
}
