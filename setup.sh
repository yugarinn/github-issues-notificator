#!/usr/bin/env sh
CompileDaemon -log-prefix=false -build "go build -o bin/github-issues-notificator ./main.go" -command "./bin/github-issues-notificator" -exclude-dir=".git"
