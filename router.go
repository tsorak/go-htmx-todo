package main

import (
	"net/http"
)

func CreateRouter() Router {
	rtr := Router{}
	rtr._routes = make(map[string]func(Responder, *http.Request))
	return rtr
}

type Router struct {
	_routes map[string]func(Responder, *http.Request)
}

func (rtr Router) addRoute(method string, path string, handler func(Responder, *http.Request)) {
	rtrKey := method + "_" + path

	rtr._routes[rtrKey] = handler
}

func (rtr Router) GET(path string, handler func(Responder, *http.Request)) {
	rtr.addRoute("GET", path, handler)
}

func (rtr Router) POST(path string, handler func(Responder, *http.Request)) {
	rtr.addRoute("POST", path, handler)
}

func (rtr Router) DELETE(path string, handler func(Responder, *http.Request)) {
	rtr.addRoute("DELETE", path, handler)
}

func (rtr Router) PATCH(path string, handler func(Responder, *http.Request)) {
	rtr.addRoute("PATCH", path, handler)
}

func (rtr Router) RouteMatcher() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		rtrKey := getRequestMethod(r) + "_" + r.URL.Path

		handler, ok := rtr._routes[rtrKey]
		if !ok {
			http.NotFound(w, r)
			return
		}
		handler(Responder{w}, r)
	}
}

func getRequestMethod(r *http.Request) string {
	method := r.Method
	if method == "" {
		return "GET"
	}
	return method
}
