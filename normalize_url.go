package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := getParsed(rawURL)
	if err != nil {
		return "", err
	}

	// Additional check for invalid URLs
	if parsedURL.Host == "" {
		return "", fmt.Errorf("invalid URL: missing host")
	}

	// Extract the host and path, ignoring the scheme
	fullPath := parsedURL.Host + parsedURL.Path
	fullPath = strings.ToLower(fullPath)
	fullPath = strings.TrimSuffix(fullPath, "/")

	return fullPath, nil
}

func getParsed(rawURL string) (*url.URL, error) {
	// Add a default scheme if missing
	if !strings.Contains(rawURL, "://") {
		rawURL = "http://" + rawURL
	}

	// Parse the raw URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("invalid URL: %v", err)
	}
	return parsedURL, nil
}
