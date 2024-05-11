package main
import (
	"fmt"
	"os"
	"strings"
)
func read_file(file string) []string {
	data, err := os.ReadFile(file)
	if err != nil {
		fmt.Print(err)
	}
	split := strings.Split(string(data), "\n\n")
	split[0] = strings.TrimPrefix(split[0], "\n")
	return split
}
func array_2d(data []string) [][]string {
	two_d_array := make([][]string, len(data))
	for i, line := range data {
		lines := strings.Split(line, "\n")
		two_d_array[i] = make([]string, len(lines))
		copy(two_d_array[i], lines)
	}
	return two_d_array
}
func print_shapes(shape [][]string, str string) {
	i := 0
	var to_print []string
	var chck []string
	chck = append(chck, str)
	if strings.Contains(str, "\\n") == true {
		chck = strings.Split(str, "\\n")
	}
	for _, line := range chck {
		for i < 8 {
			for _, char := range line {
				to_print = append(to_print, shape[int(char)-32][i])
			}
			st := strings.Join(to_print, "")
			to_print = nil
			fmt.Println(st)
			i++
		}
		i = 0
	}
}
func check_input(str string) {
	i := 0
	for range str {
		if (str[i] < 32 || str[i] > 126) && str[i] != '\n' {
			fmt.Printf("Error : Please provide a string with supported charachters !")
			os.Exit(1)
		}
		i++
	}
}
func main() {
	if len(os.Args) == 2 {
		check_input(os.Args[1])
		data := read_file("standard.txt")
		shape := array_2d(data)
		print_shapes(shape, os.Args[1])
	}
}
