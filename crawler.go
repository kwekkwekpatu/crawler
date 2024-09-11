package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func (cfg *config) addPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	_, ok := cfg.pages[normalizedURL]
	if ok {
		cfg.pages[normalizedURL]++
		return false
	}
	// Else create new entry with count of 1
	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) checkPageLimit() (pageLimitReached bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	pageLimitReached = len(cfg.pages) >= cfg.maxPages
	return pageLimitReached
}

func (cfg *config) crawlPage(rawCurrentURL string, depth int) {
	// Acquire semaphore
	cfg.concurrencyControl <- struct{}{}
	// Release semaphore when done
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()
	if cfg.checkPageLimit() {
		return
	}
	//compare base urls of both URLS
	rawBaseURL := cfg.baseURL.String()
	parsedBase, err := getParsed(rawBaseURL)
	if err != nil {
		fmt.Printf("Cannot parse rawBaseURL: %s Error %s\n", rawBaseURL, err)
		return
	}

	parsedCurrent, err := getParsed(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot parse rawCurrentURL: %s Error %s\n", rawCurrentURL, err)
		return
	}

	if parsedBase.Hostname() != parsedCurrent.Hostname() {
		fmt.Printf("Hold on wait a minute something aint right %s aint in the same domain as %s\n", parsedCurrent.Host, parsedBase.Host)
		return
	}

	// Get a normalized version of the rawCurrentURL
	normalized, err := normalizeURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot normalize URL: %s Error: %s\n", rawCurrentURL, err)
		return
	}

	// If pages already contains normalized raw count++

	if !cfg.addPageVisit(normalized) {
		fmt.Printf("Page: %s Has already been visited.\n", rawCurrentURL)
		return
	}

	// Get HTML from current URL
	resHTML, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("Cannot get HTML from: %s Error: %s\n", rawCurrentURL, err)
		return
	}
	fmt.Printf("Now printing: %s\n", rawCurrentURL)

	// Recursive call for each url in body.
	if depth >= cfg.depthLimit {
		fmt.Println("That's far enough. DepthLimit reached.")
		return
	}

	urls, err := getURLsFromHTML(resHTML, rawBaseURL)
	if err != nil {
		fmt.Printf("Cannot get URLS from HTML at: %s Error: %s\n", rawCurrentURL, err)
		return
	}
	for _, link := range urls {
		uniqueLink := link
		cfg.mu.Lock()
		cfg.wg.Add(1)
		cfg.mu.Unlock()
		go cfg.crawlPage(uniqueLink, depth+1)
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
