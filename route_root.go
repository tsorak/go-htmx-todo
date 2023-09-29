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

	fmt.Fprintf(w, "Hello world!")
}
