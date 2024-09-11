package main

import (
	"fmt"
	"os"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}
	if len(argsWithoutProg) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}
	fmt.Printf("starting crawl of: %s", argsWithoutProg[0])
	pages := make(map[string]int)
	crawlPage(argsWithoutProg[0], argsWithoutProg[0], pages)
	os.Exit(0)
}
