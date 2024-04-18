package main

import (
	"fmt"
	"os"
	"strings"
)

func ReadFile(name string) (string, error) {

	data, err := os.ReadFile(name)
	if err != nil {
	}
	return string(data), nil
}

func SplitWhiteSpaces(s string) []string {
	var result []string
	var word string
	for i, c := range s {
		if c != ' ' && c != '\t' && c != '\n' {
			word += string(c)
		}
		if ((c == ' ' || c == '\t' || c == '\n') && word != "") || i == len(s)-1 {
			result = append(result, word)
			word = ""
		}
	}
	return result
}
func ToUpper_(arr []string) []string {
	for i := 1; i < len(arr); i++ {
		if arr[i] == "(up)" && i > 0 {
			arr[i-1] = strings.ToUpper(arr[i-1])
		}
	}
	return arr
}

func ToLower_(arr []string) []string {
	for i := 1; i < len(arr); i++ {
		if arr[i] == "(low)" && i > 0 {
			arr[i-1] = strings.ToLower(arr[i-1])
		}
	}
	return arr
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("file name is not provided")
		return
	}
	file, err := ReadFile(os.Args[1])
	if err != nil {
		fmt.Printf("error reading file : %s\n", err)
		return
	}
	arr := SplitWhiteSpaces(file)
	tmp := ToLower_(arr)
	tmp = ToUpper_(arr)
	// ToUpper := ToUpper_(arr)
	// ToLower := ToLower_(arr)

	// fmt.Print(ToUpper, "\n")
	fmt.Printf(tmp, "\n")
}
