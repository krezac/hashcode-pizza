package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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
