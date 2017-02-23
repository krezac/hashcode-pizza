package main

func (cc *cacheContent) canVideoBeCached(i int, videoSizes []int, cacheSize int) bool {
	// check if the video is there
	for _, v := range cc.videos {
		if i == v {
			return false
		}
	}
	spaceUsed := 0
	for _, s := range cc.videos {
		spaceUsed += videoSizes[s]
	}
	return spaceUsed+videoSizes[i] <= cacheSize
}

func process(videoSizes []int, endpoints []endpoint, requests []request, cacheSize int, cacheContents []cacheContent) error {

	cacheContents[0].videos = []int{1, 2, 3}

	for _, r := range requests {
		for _, cl := range endpoints[r.endpoint].cacheLatencies {
			cc := &cacheContents[cl.cache]
			if cc.canVideoBeCached(r.video, videoSizes, cacheSize) {
				cc.videos = append(cc.videos, r.video)
			}
		}
	}

	return nil
}
