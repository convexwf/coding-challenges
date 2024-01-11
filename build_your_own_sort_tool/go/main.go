package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	inputFile := flag.String("file", "", "Input file to sort")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Please provide input file using -file flag")
		os.Exit(1)
	}

	// Open the input file
	file, err := os.Open(*inputFile)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read lines from file
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	// Sort the lines
	sort.Strings(lines)

	// Print the sorted lines
	for _, line := range lines {
		fmt.Println(line)
	}
}
