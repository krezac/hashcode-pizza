package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := "data/example.in"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	videoSizes, endpoints, requests, cacheSize, cacheContents, err := readFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("cache size: %d\n", cacheSize)
	fmt.Printf("video sizes (%d): %v\n", len(videoSizes), videoSizes)
	fmt.Printf("endpoints (%d): %v\n", len(endpoints), endpoints)
	fmt.Printf("requests (%d): %v\n", len(requests), requests)
	fmt.Printf("caches (%d): %v\n", len(cacheContents), cacheContents)

	err = process(videoSizes, endpoints, requests, cacheSize, cacheContents)
	if err != nil {
		panic(err)
	}
	// dump output
	f, _ := os.Create(fileName + ".out")
	defer f.Close()
	writeOutput(f, cacheContents)
}
