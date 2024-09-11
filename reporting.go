package main

import (
	"cmp"
	"fmt"
	"slices"
)

type pageResult struct {
	normalizedURL string
	count         int
}

func printReport(pages map[string]int, baseURL string) {
	fmt.Println("=============================")
	fmt.Printf("  REPORT for %s\n", baseURL)
	fmt.Println("=============================")

	sortedPages := sortPages(pages)
	for _, pageRes := range sortedPages {
		fmt.Printf("Found %v internal links to %s\n", pageRes.count, pageRes.normalizedURL)
	}
}

func sortPages(pages map[string]int) []pageResult {
	results := []pageResult{}
	for key, pageCount := range pages {
		results = append(results, pageResult{normalizedURL: key, count: pageCount})
	}
	slices.SortFunc(results, func(a, b pageResult) int {
		return cmp.Or(
			cmp.Compare(a.count, b.count),
			cmp.Compare(a.normalizedURL, b.normalizedURL),
		)
	})
	return results
}
