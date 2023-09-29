package main

import (
	"fmt"
	"net/http"
)

func routeRoot(w http.ResponseWriter, r *http.Request) {
	if !accepts("text/html", r) {
		w.WriteHeader(http.StatusUnavailableForLegalReasons)
		return
	}

	htmx, _ := getFileContents("htmx/root.html")

	fmt.Fprint(w, htmx)
}
