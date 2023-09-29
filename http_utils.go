package main

import (
	"net/http"
	"strings"
)

func accepts(pattern string, r *http.Request) bool {
	accepts := r.Header.Get("Accept")
	return strings.Contains(accepts, pattern)
}
