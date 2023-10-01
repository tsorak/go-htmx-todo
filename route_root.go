package main

import (
	"net/http"
)

func routeRoot(resp Responder, r *http.Request) {
	if !accepts("text/html", r) {
		resp.SetStatus(http.StatusNotAcceptable)
		return
	}

	htmx, _ := getFileContents("htmx/root.html")

	resp.SendBytes(htmx)
}
