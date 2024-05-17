package main

import (
	"fmt"
	"os"
	"strings"
)

func contains(input string, c rune) bool {
	return strings.ContainsRune(input, c)
}

func print_shapes(shape [][]string, str string) {
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run . --color=<color> <substring> <string>")
		return
	}
	color := os.Args[2][8:] // Extract color from --color=<color>
	input := os.Args[3]

	i := 0
	var to_print []string
	var new []string
	str = strings.ReplaceAll(str, "\\n", "\n")
	if strings.Contains(str, "\n") {
		new = strings.Split(str, "\n")
	} else {
		new = append(new, str)
	}

	colorCode := getColorCode(color)

	for _, line := range new {
		if line != "" {
			for i < 8 {
				for _, c := range line {
					if contains(input, c) {
						// Print the character in the specified color
						to_print = append(to_print, fmt.Sprintf("%s%s%s", colorCode, shape[int(c)-32][i], "\033[0m"))
					} else {
						to_print = append(to_print, shape[int(c)-32][i])
					}
				}
				st := strings.Join(to_print, "")
				to_print = nil
				fmt.Printf("%s\n", st)
				i++
			}
			i = 0
		} else {
			fmt.Printf("\n")
		}
	}
}

func getColorCode(color string) string {
	switch strings.ToLower(color) {
	case "red":
		return "\033[31m"
	case "green":
		return "\033[32m"
	case "yellow":
		return "\033[33m"
	case "blue":
		return "\033[34m"
	case "magenta":
		return "\033[35m"
	case "cyan":
		return "\033[36m"
	case "white":
		return "\033[37m"
	default:
		return "\033[0m" // Default color (reset)
	}
}

func main() {
	// Example shape definition for ASCII art (simplified)
	shape := [][]string{
		{" ", " ", " ", " ", " ", " ", " ", " "}, // space
		{"H", "H", " ", " ", "H", "H", "H", "H"}, // H
		{"E", "E", "E", "E", " ", " ", " ", " "}, // E
		{"L", " ", " ", " ", " ", " ", " ", " "}, // L
		{"L", " ", " ", " ", " ", " ", " ", " "}, // L
		{"O", "O", "O", "O", " ", " ", " ", " "}, // O
		// Add more characters as needed
	}

	str := os.Args[4]
	print_shapes(shape, str)
}
