package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// Validate The os.Args input
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . [string]\n	OR")
		fmt.Println("Usage: go run . [string] standard/thinkertoy/shadow")

		os.Exit(1)
	}

	// Get Input
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
		fmt.Printf("the banner you gave %q is not a among valid banners standard/thinkertoy/shadow \n", bannerName)
		os.Exit(1)
	}

	//SplitInput
	lines := splitInput(input)

	//loadBanner
	banner := loadBanner(bannerName + ".txt")

	//printArt
	printArt(lines, banner)

}

func loadBanner(filename string) []string {
	texts, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: Error reading file", err)
		return nil
	}
	result := strings.Split(string(texts), "\n")
	return result
}

func splitInput(text string) []string {
	texts := strings.Split(text, "\\n")
	return texts
}

func printArt(lines []string, banner []string) {
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}
		for row := 0; row < 9; row++ {
			for _, ch := range line {
				ascii := int(ch)
				fmt.Print(banner[(ascii-32)*9+row])

			}
			fmt.Println()
		}
	}
}
