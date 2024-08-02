package main

import (
	"fmt"
	"net/url"
	"strings"
)

func NormalizeURL(rawURL string) (string, error) {
	// Add a default scheme if missing
	if !strings.Contains(rawURL, "://") {
		rawURL = "http://" + rawURL
	}

	// Parse the raw URL
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %v", err)
	}

	// Additional check for invalid URLs
	if parsedURL.Host == "" {
		return "", fmt.Errorf("invalid URL: missing host")
	}

	// Extract the host and path, ignoring the scheme
	normalizedURL := strings.TrimSuffix(parsedURL.Host+parsedURL.Path, "/")
	return normalizedURL, nil
}
