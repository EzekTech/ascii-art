package main

/*import (
	"fmt"
	"os"
	"strings"
)*/

/*func main() {
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
*/
/*
func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . text\n		OR")
		fmt.Println("Usage: go run . text standard/shadow/thinkertoy")
	}

	input := os.Args[1]

	lines := splitInput(input)

	bannerName := "standard"
	if len(os.Args) == 3 {
		bannerName = os.Args[2]
	}

	validBanner := map[string]bool{
		"standard":   true,
		"thinkertoy": true,
		"shadow":     true,
	}

	banner := loadBanner(bannerName + ".txt")

	if !validBanner[bannerName+".txt"] {
		fmt.Printf("The banner name  %q is not a valid banner, valid banners are standard/thinkertoy/shadow", bannerName)
	}

	printArt(lines, banner)

}

func loadBanner(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Errorf("Error reading banner file %q : %w", fileName, err)
		os.Exit(1)
	}
	result := strings.Split(string(data), "\n")
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
*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: go run . text\n		OR")
		fmt.Println("Usage: go run . text standard/shadow/thinkertoy")
		os.Exit(1)
	}

	input := os.Args[1]
	lines := strings.Split(input, "\\n")

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
		fmt.Printf("The banner you gave %q is not among the valid banners standard/thinkertoy/shaadow\n", bannerName)
		return
	}

	banner := loadBanner(bannerName + ".txt")

	printArt(lines, banner)
}

func loadBanner(fileName string) []string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading banner file")
	}

	content := strings.ReplaceAll(string(data), "\r\n", "\n")
	result := strings.Split(content, "\n")
	return result
}

/*func splitInput(text string) []string {
	texts := strings.Split(text, "\\n")
	//convert := strings.Replace(string(text), "\r\n", "\n")
	return texts
}*/

func printArt(lines []string, banner []string) {
	for _, line := range lines {
		if line == "" {
			fmt.Println()
			continue
		}

		for row := 1; row < 9; row++ {
			for _, ch := range line {
				ascii := int(ch)
				if ascii < 32 || ascii > 126 {
					continue
				}
				fmt.Print(banner[(ascii-32)*9+row])

			}
			fmt.Println()
		}
		//fmt.Println()
	}
}
