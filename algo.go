package main

func (cc *cacheContent) canVideoBeCached(i int, videoSizes []int, cacheSize int) bool {
	spaceUsed := 0
	for _, s := range cc.videos {
		spaceUsed += videoSizes[s]
	}
	return spaceUsed+videoSizes[i] <= cacheSize
}

func process(videoSizes []int, endpoints []endpoint, requests []request, cacheSize int, cacheContents []cacheContent) error {

	// TODO remove - test data
	cacheContents[4].videos = []int{4, 5, 6}

	return nil
}
