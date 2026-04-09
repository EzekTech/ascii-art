package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Validate arguments
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [string]\n\tOR")
		fmt.Println("Usage: go run . [string] standard/thinkertoy/shadow")
		os.Exit(1)
	}

	// Get input
	input := os.Args[1]

	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	validBanner := map[string]bool{
		"standard":   true,
		"thinkertoy": true,
		"shadow":     true,
	}

	if !validBanner[bannerName] {
		fmt.Printf("the banner you gave %q is not among valid banners standard/thinkertoy/shadow\n", bannerName)
		os.Exit(1)
	}

	// Normalize input (VERY IMPORTANT)
	input = strings.ReplaceAll(input, "\\n", "\n")

	// Split input into lines
	lines := splitInput(input)

	// Load banner
	banner := loadBanner(bannerName + ".txt")

	// Print ASCII art
	printArt(lines, banner)
}

func loadBanner(filename string) []string {
	texts, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: Error reading file", err)
		os.Exit(1)
	}
	return strings.Split(string(texts), "\n")
}

func splitInput(text string) []string {
	return strings.Split(text, "\n")
}

func printArt(lines []string, banner []string) {
	for i, line := range lines {

		// Prevent leading empty line
		if i == 0 && line == "" {
			continue
		}

		// Handle empty line (just print ONE newline)
		if line == "" {
			fmt.Println()
			continue
		}

		// Each ASCII char is 8 rows tall
		for row := 1; row < 9; row++ {
			for _, ch := range line {
				ascii := int(ch)

				// Safety check (print nothing if out of range)
				if ascii < 32 || ascii > 126 {
					continue
				}

				fmt.Print(banner[(ascii-32)*9+row])
			}
			fmt.Println()
		}
	}
}
