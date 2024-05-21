package main

import (
	"fmt"
	"os"
	"strings"
)

func read_file(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print("Error: reading the file\n")
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

func GetColor(color string) string {
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
	default:
		return "\033[0m"
	}
}

// func Matched(str, substr, colorCode string) string {
// 	var result strings.Builder
// 	for _, c := range str {
// 		if strings.ContainsRune(substr, c) {
// 			result.WriteString(fmt.Sprintf("%s%c%s", colorCode, c, "\033[0m"))
// 		} else {
// 			result.WriteRune(c)
// 		}
// 	}
// 	return result.String()
// }

func getIndex(line, input string) []int {
	indx := []int{}
	for index1 := 0; len(line)-len(input) >= index1; index1++ {
		if line[index1:index1+len(input)] == input {
			indx = append(indx, index1)
		}
	}
	return indx
}

func print_shapes(shape [][]string, str string, color string, input string) {
	i := 0
	// tmp := ""
	var to_print []string
	var new []string
	str = strings.ReplaceAll(str, "\\n", "\n")
	if strings.Contains(str, "\n") {
		new = strings.Split(str, "\n")
	} else {
		new = append(new, str)
	}
	colorCode := GetColor(color)
	for _, line := range new {
		if line != "" {
			for i < 8 {
				for index1, c := range line {
					if strings.Contains(line, input) {
						count := getIndex(line, input)
						// fmt.Print(count)
						// os.Exit(1)
						for j, n := range count {
							if index1 >= n && index1 <= n+len(input)-1 {
								to_print = append(to_print, fmt.Sprintf("%s%s%s", colorCode, shape[int(c)-32][i], "\033[0m"))
								break
							}
							if j == len(count)-1 && !(index1 >= n && index1 <= n+len(input)-1) {
								to_print = append(to_print, shape[int(c)-32][i])
							}
						}
						// to_print = append(to_print, shape[int(c)-32][i])

					} else {
						to_print = append(to_print, shape[int(c)-32][i]) // replace the indexes i had that need to get colored with the colored ones in the loop if it functions
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
			fmt.Printf("\033[1m\033[31mError: Please provide a string with supported characters!\n")
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

func checkflag(str string) string {
	if strings.HasPrefix(str, "--color=") {
		return str[8:]
	}
	if strings.HasPrefix(str, "--output=") {
		return str[8:]
	}
	return ""
}

func main() {
	if len(os.Args) >= 2 && len(os.Args) <= 4 {
		if check_input(os.Args[1]) {
			return
		}
		if check_newline(os.Args[1]) {
			return
		}
		data := read_file("standard.txt")
		shape := array_2d(data)
		// color := os.Args[1][8:]
		color := checkflag(os.Args[1])
		input := os.Args[2]
		print_shapes(shape, os.Args[3], color, input)
	} else {
		fmt.Printf("\033[1m\033[033mError: enter the correct arguments!\n")
	}
}

