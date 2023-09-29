package main

import (
	"net/http"
)

func startWebserver() {
	logServerAdress()
	http.HandleFunc("/", router)
	http.ListenAndServe(":8080", nil)
}
