package main

import (
	"net/http"
)

func routeMotd(resp Responder, r *http.Request) {
	htmx, _ := getFileContents("htmx/motd.html")

	resp.SendBytes(htmx)
}
