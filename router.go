package main

import (
	"fmt"
	"net/http"
)

func router(w http.ResponseWriter, r *http.Request) {
	var route string = r.URL.Path

	fmt.Println(route)

	switch route {
	case "/":
		routeRoot(w, r)
	case "/motd":
		routeMotd(w, r)
	default:
		routeNotFound(w, r)
	}
}

func routeNotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
