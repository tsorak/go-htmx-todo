package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func accepts(pattern string, r *http.Request) bool {
	accepts := r.Header.Get("Accept")
	return strings.Contains(accepts, pattern) || strings.Contains(accepts, "*/*")
}

func getFileContents(path string) ([]byte, error) {
	absPath, _ := filepath.Abs(path)
	fsfile, err := os.Open(absPath)
	if err != nil {
		return make([]byte, 0), err
	}

	bytes := make([]byte, 1024)
	fsfile.Read(bytes)

	return bytes, nil
}

func getFileContentsIntoString(path string) (string, error) {
	absPath, _ := filepath.Abs(path)
	bytes, err := getFileContents(absPath)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
