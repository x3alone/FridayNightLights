package main
import (
	"fmt"
	"os"
	"strings"
)
func read_file(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print("Error : reading the file\n")
		os.Exit(1)
	}
	split := strings.Split(string(data), "\n\n")
	split[0] = strings.TrimPrefix(split[0], "\n")
	return split
}
func array_2d(data []string) [][]string { 
	two_d_array := make([][]string, len(data))
	for i, line := range data {
		lines := strings.Split(line, "\n")
		two_d_array[i] = lines
	}
	return two_d_array
}
func contains(input string, c rune) bool {
	return strings.ContainsRune(input, c)
}
func print_shapes(shape [][]string, str string) {
	color := os.Args[2][8:] 
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
func check_input(str string) bool {
	i := 0
	for range str {
		if (str[i] < 32 || str[i] > 126) && str[i] != '\n' {
			fmt.Printf("\033[1m\033[31mError : Please provide a string with supported charachters!\n")
			return true
		}
		i++
	}
	return false
}
func all_newline(str string) bool {
	i := 0
	for range str {
		if str[i] == '\n' {
			i++
		}
	}
	return len(str) == i
}
func check_newline(str string) bool {
	str = strings.ReplaceAll(str, "\\n", "\n")
	if len(str) == 0 {
		return true
	} else if len(str) == 1 && str[0] == '\n' {
		fmt.Print("\n")
		return true
	} else if all_newline(str) {
		for i := 0; len(str) > i; i++ {
			fmt.Print("\n")
		}
		return true
	}
	return false
}
func main() {
	if len(os.Args) > 2 && len(os.Args) < 4 {
		if check_input(os.Args[1]) {
			return
		}
		if check_newline(os.Args[1]) {
			return
		}
		data := read_file("standard.txt")
		shape := array_2d(data)
		print_shapes(shape, os.Args[1])
	} else {
		fmt.Printf("\033[1m\033[033mError: enter the correct arguments!\n")
	}
}