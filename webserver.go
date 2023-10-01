package main

import (
	"net/http"
)

func startWebserver() {
	logServerAdress()

	rtr := CreateRouter()
	setupRoutes(&rtr)

	http.HandleFunc("/", rtr.RouteMatcher())
	http.ListenAndServe(":8080", nil)
}

func setupRoutes(rtr *Router) {
	rtr.GET("/", routeRoot)
	rtr.GET("/motd", routeMotd)

	rtr.POST("/todo", routeTodoAdd)
	rtr.DELETE("/todo", routeTodoDelete)
}
