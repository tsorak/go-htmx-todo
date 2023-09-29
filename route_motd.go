package main

import (
	"fmt"
	"net/http"
)

func routeMotd(w http.ResponseWriter, r *http.Request) {
	htmx, _ := getFileContents("htmx/motd.html")

	fmt.Fprint(w, htmx)
}
