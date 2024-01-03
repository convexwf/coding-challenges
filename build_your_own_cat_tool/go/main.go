package main

import (
	"flag"
	"fmt"
	"os"
)

func cat(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		return
	}
	fmt.Printf("%s", data)
}

func main() {
	lineNumbers := flag.Bool("n", false, "Display line numbers")
	flag.Parse()

	if flag.NArg() == 0 {
		fmt.Fprintln(os.Stderr, "Usage: cat [options] file1 file2 ...")
		os.Exit(1)
	}

	for _, filename := range flag.Args() {
		cat(filename)
	}

	if *lineNumbers {
		fmt.Println()
	}
}
