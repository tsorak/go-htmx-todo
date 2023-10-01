package main

import "net/http"

func routeTodoDelete(resp Responder, r *http.Request) {
	targetId := r.URL.Query().Get("id")

	success := db.deleteTodoByID(targetId)

	if !success {
		resp.SetStatus(http.StatusBadRequest)
		return
	}

	htmx := ""
	resp.SetHeader("Content-Type", "text/html").SetStatus(http.StatusOK).SendString(htmx)
}
