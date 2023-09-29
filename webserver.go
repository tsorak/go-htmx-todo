package main

import (
	"net/http"
)

func startWebserver() {
	logServerAdress()
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}

func setupRoutes() {
	http.HandleFunc("/", routeRoot)
}
