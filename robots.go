package main

import (
	"fmt"
	"io"
	"net/http"
)

func fetchRobotsTXT(baseURL string) (string, error) {
	robotsURL := baseURL + "/robots.txt"
	resp, err := http.Get(robotsURL)
	if err != nil {
		return "", fmt.Errorf("failed to fetch robots.txt: %w", err)
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("Received error with status: %s", resp.Status)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read robots.txt body: %w", err)
	}

	return string(body), nil

}

func parseRobotsTXT(robotsString string) {

}
