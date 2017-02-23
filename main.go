package main

import "os"

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

	slices := process(pizza, minItems, maxSize)

	// dump output
	f, _ := os.Create(fileName + ".out")
	defer f.Close()
	writeOutput(f, slices)
}
