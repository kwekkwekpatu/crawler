package main

import (
	"fmt"
	"io"
	"net/http"
)

func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("Received error with status: %s", resp.Status)
	}

	if resp.Header.Get("content-type") != "text/html" {
		return "", fmt.Errorf("Error content-type is not text/html but %s", resp.Header.Get("content-type"))
	}

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading HTML body")
	}

	return string(page), nil
}
