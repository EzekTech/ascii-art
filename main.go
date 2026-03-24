package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// Check if the user passed exactly one argument
	if len(os.Args) != 2 {
		fmt.Println("usage : go run . \"text\"")
		return
	}

	// Store the user input from the command line
	input := os.Args[1]

	// Read the ASCII banner file ("standard.txt")
	data, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error Reading file", err)
	}

	// Convert the file content into line. Each line represents part of the ASCII characters
	line := strings.Split(string(data), "\n")

	// Call the function that converts the text into ASCII art
	result := printASCII(input, line)
	// Print the final ASCII result to the terminal 
	fmt.Print(result)

}

func printASCII(input string, line []string) string {
	if strings.ReplaceAll(input, "\\n", "") == "" {
		count := len(input) / 2
		return strings.Repeat("\n", count)
	}

	var output strings.Builder
	userInput := strings.Split(input, "\\n")

	for _, word := range userInput {
		if word != "" {
			for row := 0; row < 8; row++ {
				for _, char := range word {
					index := (int(char)-32)*9 + 1 + row
					output.WriteString(line[index])
				}

				output.WriteString("\n")
			}
		} else {
			output.WriteString("\n")
		}
	}

	return output.String()
}
