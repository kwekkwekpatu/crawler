package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	maxConcurrency, err := strconv.Atoi(argsWithoutProg[1])
	if err != nil {
		fmt.Printf("Invalid second argument: %v reverting to maxConcurrency default of 5", err)
		maxConcurrency = 5
	}

	maxPages, err := strconv.Atoi(argsWithoutProg[2])
	if err != nil {
		fmt.Printf("Invalid second argument: %v reverting to maxPages default of 50", err)
		maxConcurrency = 5
	}
	const depthLimit = 50

	cfg, err := configure(argsWithoutProg[0], maxConcurrency, depthLimit, maxPages)
	if err != nil {
		fmt.Printf("Error configuring: %v", err)
	}

	fmt.Printf("starting crawl of: %s", argsWithoutProg[0])
	cfg.mu.Lock()
	cfg.wg.Add(1)
	cfg.mu.Unlock()
	go cfg.crawlPage(argsWithoutProg[0], 0)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}

	os.Exit(0)
}
