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
	result, err := getHTML(argsWithoutProg[0])
	if err != nil {
		fmt.Printf("Error reading HTML body: %s", err)
		os.Exit(1)
	}
	fmt.Println(result)
	os.Exit(0)
}
