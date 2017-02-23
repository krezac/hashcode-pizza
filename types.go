package main

type cacheLatency struct {
	cache   int
	latency int
}

type endpoint struct {
	latency        int            // endpoint latency
	cacheLatencies []cacheLatency // caches latencies
}

type request struct {
	video    int
	endpoint int
	count    int
}

type cacheContent struct {
	videos []int
}
