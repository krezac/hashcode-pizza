package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readFile(fileName string) ([]int, []endpoint, []request, int, []cacheContent, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, nil, nil, 0, nil, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	// read header
	scanner.Scan()
	headers := strings.Split(scanner.Text(), " ")
	// and parse it
	videoCount, err := strconv.Atoi(headers[0])
	if err != nil {
		return nil, nil, nil, 0, nil, err
	}
	endpointCount, err := strconv.Atoi(headers[1])
	if err != nil {
		return nil, nil, nil, 0, nil, err
	}
	requestCount, err := strconv.Atoi(headers[2])
	if err != nil {
		return nil, nil, nil, 0, nil, err
	}
	cacheCount, err := strconv.Atoi(headers[3])
	if err != nil {
		return nil, nil, nil, 0, nil, err
	}
	cacheSize, err := strconv.Atoi(headers[4])
	if err != nil {
		return nil, nil, nil, 0, nil, err
	}

	videoSizes := make([]int, videoCount)
	scanner.Scan()
	videosStr := strings.Split(scanner.Text(), " ")
	if videoCount != len(videosStr) {
		return nil, nil, nil, 0, nil, fmt.Errorf("incorrect video sizes count")
	}
	for i, v := range videosStr {
		vi, err := strconv.Atoi(v)
		if err != nil {
			return nil, nil, nil, 0, nil, err
		}
		videoSizes[i] = vi
	}

	endpoints := make([]endpoint, endpointCount)
	for i := 0; i < endpointCount; i++ {
		scanner.Scan()
		e := strings.Split(scanner.Text(), " ")
		el, err := strconv.Atoi(e[0])
		if err != nil {
			return nil, nil, nil, 0, nil, err
		}
		cc, err := strconv.Atoi(e[1])
		if err != nil {
			return nil, nil, nil, 0, nil, err
		}

		// iterate the latencies
		latencies := make([]cacheLatency, cc)
		for i := 0; i < cc; i++ {
			scanner.Scan()
			e := strings.Split(scanner.Text(), " ")
			n, err := strconv.Atoi(e[0])
			if err != nil {
				return nil, nil, nil, 0, nil, err
			}
			if n >= cacheCount {
				return nil, nil, nil, 0, nil, fmt.Errorf("incorrect cache count")
			}
			lat, err := strconv.Atoi(e[1])
			if err != nil {
				return nil, nil, nil, 0, nil, err
			}
			latencies[i].cache = n
			latencies[i].latency = lat
		}

		endpoints[i].latency = el
		endpoints[i].cacheLatencies = latencies
	}

	requests := make([]request, requestCount)
	// read requests
	for i := 0; i < requestCount; i++ {
		scanner.Scan()
		e := strings.Split(scanner.Text(), " ")
		video, err := strconv.Atoi(e[0])
		if err != nil {
			return nil, nil, nil, 0, nil, err
		}
		endpoint, err := strconv.Atoi(e[1])
		if err != nil {
			return nil, nil, nil, 0, nil, err
		}
		reqCount, err := strconv.Atoi(e[2])
		if err != nil {
			return nil, nil, nil, 0, nil, err
		}

		requests[i].video = video
		requests[i].endpoint = endpoint
		requests[i].count = reqCount
	}

	cacheContents := make([]cacheContent, cacheCount)

	return videoSizes, endpoints, requests, cacheSize, cacheContents, nil
}
