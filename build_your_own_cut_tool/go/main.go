package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	delimiter := flag.String("d", " ", "set the field delimiter")
	fields := flag.String("f", "1", "select only these fields")
	flag.Parse()

	// Check if filename is provided as argument
	if flag.NArg() == 0 {
		fmt.Println("No input file provided")
		os.Exit(1)
	}

	// Open input file
	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fieldsList := strings.Split(line, *delimiter)
		for _, fieldStr := range strings.Split(*fields, ",") {
			fieldNum, err := strconv.Atoi(fieldStr)
			if err != nil {
				fmt.Println("Invalid field number:", fieldStr)
				os.Exit(1)
			}

			if fieldNum <= len(fieldsList) {
				fmt.Print(fieldsList[fieldNum-1])
				if fieldNum < len(fieldsList) {
					fmt.Print(*delimiter)
				}
			}
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input file:", err)
		os.Exit(1)
	}
}
