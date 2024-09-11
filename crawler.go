package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	//compare base urls of both URLS
	parsedBase, err := getParsed(rawBaseURL)
	if err != nil {
		fmt.Printf("Cannot parse rawBaseURL: %s Error %s\n", rawBaseURL, err)
	}

	parsedCurrent, err := getParsed(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot parse rawCurrentURL: %s Error %s\n", rawCurrentURL, err)
	}

	if parsedBase.Host != parsedCurrent.Host {
		fmt.Printf("Hold on wait a minute something aint right %s aint in the same domain as %s\n", parsedCurrent.Host, parsedBase.Host)
		return
	}
	fmt.Println("1) Parsed the URLS")

	// Get a normalized version of the rawCurrentURL
	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot normalize URL: %s Error: %s\n", rawCurrentURL, err)
		return
	}

	fmt.Println("2) Normalized the currentURL")
	// If pages already contains normalized raw count++
	_, ok := pages[normalized]
	if ok {
		pages[normalized] = pages[normalized] + 1
		fmt.Println("3a) Page has been visited before skipping")
		return
	}
	// Else create new entry with count of 1
	pages[normalized] = 1
	fmt.Println("3b) New entry created")
	// Get HTML from current URL
	resHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot get HTML from: %s Error: %s\n", rawCurrentURL, err)
		return
	}
	fmt.Println("4) Got the HTML from the URL")
	fmt.Printf("Now printing: %s\n", rawCurrentURL)
	fmt.Println(resHTML)
	// Recursive call for each url in body.
	urls, err := getURLsFromHTML(resHTML, rawBaseURL)
	fmt.Println("5) Got the URLS from the HTML")
	for _, link := range urls {
		crawlPage(rawBaseURL, link, pages)
	}
}
func getHTML(rawURL string) (string, error) {
	resp, err := http.Get(rawURL)
	if err != nil {
		return "", err
	}
	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("Received error with status: %s", resp.Status)
	}

	if !strings.Contains(resp.Header.Get("content-type"), "text/html") {
		return "", fmt.Errorf("Error content-type is not text/html but %s", resp.Header.Get("content-type"))
	}

	page, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Error reading HTML body")
	}

	return string(page), nil
}
