package http

import (
	"log"
	"net/http"

	"github.com/yugarinn/github-issues-notificator/core"
)


// TODO: move this to .env
const HTTP_PORT = ":8081"

func InitServer(app *core.App) {
	http.HandleFunc("/notifications", createNotificationHandler)

	log.Fatal(http.ListenAndServe(HTTP_PORT, nil))
}
