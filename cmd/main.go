package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)

func main() {
	// Define command-line flags
	csvfile := flag.String("c", "", "CSV file to compare")
	row := flag.Int("r", 0, "Row Number of CSV file to compare")
	target := flag.String("t", "", "Specified Value to compare")
	listfile := flag.String("l", "", "List file to compare")
	flag.Parse()

	// Check if both files are provided
	if *csvfile == "" || *row <= 0 {
		fmt.Println("Both of csvfile and row must be provided")
		flag.Usage()
		os.Exit(1)
	}

	if *target == "" && *listfile == "" {
		fmt.Println("target or List file must be specified")
		flag.Usage()
		os.Exit(1)
	}

	// Open the CSV file
	file, err := os.Open(*csvfile)
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		os.Exit(1)
	}

	if *listfile != "" {
		// Open the list file
		lfile, err := os.Open(*listfile)
		if err != nil {
			fmt.Println("Error opening CSV file:", err)
			os.Exit(1)
		}
		defer file.Close()

		// Print the third column of each row
		for _, record := range records {
			if len(record) >= *row {
				// Read the file line by line
				scanner := bufio.NewScanner(lfile)
				for scanner.Scan() {
					if record[*row-1] == scanner.Text() {
						fmt.Println(scanner.Text())
					}
				}
			} else {
				fmt.Println("Row does not have a third column")
			}
		}
	} else {
		// Print the third column of each row
		for _, record := range records {
			if len(record) >= *row {
				if *target == record[*row-1] {
					fmt.Println("Compared Value Matched:")
					fmt.Println(*target)
				}
			} else {
				fmt.Println("Row does not have a third column")
			}
		}
	}

	// Placeholder for CSV comparison logic
	fmt.Printf("Comparing %s and %s\n", *csvfile, *listfile)
}
