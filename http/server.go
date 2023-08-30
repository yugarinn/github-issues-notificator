package http

import (
	"log"
	"net/http"
)


// TODO: move this to .env
const HTTP_PORT = ":8081"

func InitServer() {
	http.HandleFunc("/notifications", createNotificationHandler)

	log.Fatal(http.ListenAndServe(HTTP_PORT, nil))
}
