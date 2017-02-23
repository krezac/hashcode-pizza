package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const (
	Mushroom = 'M'
	Tomato   = 'T'
)

type slice struct {
	r0, r1, c0, c1 int
}

func readFile(fileName string) ([]string, int, int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	// read header
	scanner.Scan()
	headers := strings.Split(scanner.Text(), " ")
	// and parse it
	rows, err := strconv.Atoi(headers[0])
	if err != nil {
		return nil, 0, 0, err
	}
	columns, err := strconv.Atoi(headers[1])
	if err != nil {
		return nil, 0, 0, err
	}
	minItems, err := strconv.Atoi(headers[2])
	if err != nil {
		return nil, 0, 0, err
	}
	maxSize, err := strconv.Atoi(headers[3])
	if err != nil {
		return nil, 0, 0, err
	}
	pizza := []string{}
	for scanner.Scan() {
		pizza = append(pizza, scanner.Text())
	}
	// sanity check
	if len(pizza) != rows {
		return nil, 0, 0, fmt.Errorf("incorrect row count")
	}
	for _, row := range pizza {
		if len(row) != columns {
			return nil, 0, 0, fmt.Errorf("incorrect column count")
		}
	}
	return pizza, minItems, maxSize, nil
}

func (s *slice) isValid(pizza []string, minItems, maxSize int) bool {
	if (s.r1-s.r0+1)*(s.c1-s.c0+1) > maxSize {
		return false
	}
	mushrooms, tomatoes := 0, 0
	for i := s.r0; i <= s.r1; i++ {
		for j := s.c0; j <= s.c1; j++ {
			if pizza[i][j] == Tomato {
				tomatoes++
			} else if pizza[i][j] == Mushroom {
				mushrooms++
			} else {
				panic("Invalid char in pizza")
			}
		}
		if tomatoes < minItems || mushrooms < minItems {
			return false
		}
	}
	return true
}

func (s *slice) write(w io.Writer) {
	fmt.Fprintf(w, "%d %d %d %d\n", s.r0, s.c0, s.r1, s.c1)
}

func writeOutput(w io.Writer, slices []slice) {
	fmt.Fprintf(w, "%d\n", len(slices))
	for _, s := range slices {
		s.write(w)
	}
}

func main() {
	fileName := "data/example.in"
	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}
	pizza, minItems, maxSize, err := readFile(fileName)
	if err != nil {
		panic(err)
	}
	//fmt.Println(pizza, minItems, maxSize)

	slices := []slice{}

	for r, row := range pizza {
		start := 0
		for start < len(row) {
			found := false
			for i := 0; i < maxSize && start+i < len(row); i++ {
				//fmt.Printf("row %d, start: %d, len: %d\n", r, start, i)
				slice := slice{
					r0: r,
					r1: r,
					c0: start,
					c1: start + i,
				}
				if slice.isValid(pizza, minItems, maxSize) {
					slices = append(slices, slice) // test line
					start = slice.c1 + 1
					//fmt.Printf("found, new start: %d\n", start)
					found = true
					break
				}
			}
			if !found {
				start++
			}
		}
	}

	// dump output
	f, _ := os.Create(fileName + ".out")
	defer f.Close()
	writeOutput(f, slices)
}
