package main

import (
	"fmt"
	"io"
)

func (s *slice) write(w io.Writer) {
	fmt.Fprintf(w, "%d %d %d %d\n", s.r0, s.c0, s.r1, s.c1)
}

func writeOutput(w io.Writer, slices []slice) {
	fmt.Fprintf(w, "%d\n", len(slices))
	for _, s := range slices {
		s.write(w)
	}
}
