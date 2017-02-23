package main

import (
	"fmt"
	"io"
)

func usedCaches(c []cacheContent) int {
	used := 0
	for _, cc := range c {
		if len(cc.videos) > 0 {
			used++
		}
	}
	return used
}

func (c *cacheContent) write(w io.Writer, i int) {
	fmt.Fprintf(w, "%d", i)
	for _, v := range c.videos {
		fmt.Fprintf(w, " %d", v)
	}
	fmt.Fprintln(w)
}

func writeOutput(w io.Writer, out []cacheContent) {
	fmt.Fprintf(w, "%d\n", usedCaches(out))
	for i, cc := range out {
		if len(cc.videos) > 0 {
			cc.write(w, i)
		}
	}
}
