package web

import (
	"os"
	"strings"
)

/* reads the specified file from the banner*/
func read_file(file string) []string {
	data, _ := os.ReadFile("AsciiWeb/" + file)
	if file == "thinkertoy.txt" {
		data = []byte(strings.ReplaceAll(string(data), "\r", ""))
	}
	split := strings.Split(string(data), "\n\n")
	split[0] = strings.TrimPrefix(split[0], "\n")
	return split
}

/* converts the read data to a 2d array*/
func array_2d(data []string) [][]string {
	two_d_array := make([][]string, len(data))
	for i, line := range data {
		lines := strings.Split(line, "\n")
		two_d_array[i] = lines
	}
	return two_d_array
}

/* the function that prints  the shapes*/
func print_shapes(shape [][]string, str string) string {
	i := 0
	var to_print []string
	var new []string
	var result string
	str = strings.ReplaceAll(str, "\n", "\n")
	str = strings.ReplaceAll(str, "\r\n", "\n")
	 if strings.Contains(str, "\n") {
		new = strings.Split(str, "\n")
	} else {
		new = append(new, str)
	}
	for _, line := range new {
		if line != "" {
			for i < 8 {
				for _, c := range line {
					to_print = append(to_print, shape[int(c)-32][i])
				}
				st := strings.Join(to_print, "")
				to_print = nil
				result += st + "\n"
				i++
			}
			i = 0
		} else {
			result += "\n"
		}
	}
	return result
}

/*checks that the input from the terminal is comform*/
func check_input(str string) bool {
	i := 0
	str = strings.ReplaceAll(str, "\r\n", "\n")
	for range str {
		if (str[i] < 32 || str[i] > 126) && str[i] != '\n' {
			return true
		}
		i++
	}
	return false
}

/* checks if the string is all new lines */
func all_newline(str string) bool {
	i := 0
	for range str {
		if str[i] == '\n' {
			i++
		}
	}
	return len(str) == i
}

/* handles the cases of a  single new line in the string or multiple ones */
func check_newline(str string) (string, bool) {
	var result string
	str = strings.ReplaceAll(str, "\\n", "\n")
	if len(str) == 0 {
		return "", true
	} else if len(str) == 1 && str[0] == '\n' {
		return "\n", true
	} else if all_newline(str) {
		for i := 0; len(str) > i; i++ {
			result += "\n"
		}
		return result, true
	}
	return "", false
}
