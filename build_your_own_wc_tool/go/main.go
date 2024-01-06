package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wc <filename>")
		return
	}

	filename := os.Args[1]

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := 0
	words := 0
	chars := 0

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		words += len(strings.Fields(line))
		chars += len(line)
	}

	fmt.Printf("Lines: %d\n", lines)
	fmt.Printf("Words: %d\n", words)
	fmt.Printf("Characters: %d\n", chars)
}
